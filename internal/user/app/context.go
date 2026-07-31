package app

import (
	"context"
	"time"

	medicalRepo "github.com/teakingwang/backend/internal/user/medical/repository"
	medicalSvc "github.com/teakingwang/backend/internal/user/medical/service"
	"github.com/teakingwang/backend/internal/user/model"
	"gi
	"github.com/teakingwang/backend/internal/user/repository"
	"github.com/teakingwang/backend/internal/user/service"
	"github.com/teakingwang/backend/pkg/datastore"
	"github.com/teakingwang/backend/pkg/db"
	"github.com/teakingwang/backend/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AppContext struct {
	DB                   *gorm.DB
	Redis                datastore.Store
	UserService          service.UserService
	RBACService          service.RBACService
	MedicalRecordService medicalSvc.MedicalRecordService
}

// nicknameProvider 通过 UserService（未来可改为 gRPC 客户端）获取用户昵称
type nicknameProvider struct {
	userSvc service.UserService
}

func (p *nicknameProvider) GetUserNickname(ctx context.Context, userID int64) string {
	dto, err := p.userSvc.GetUser(ctx, userID)
	if err != nil || dto == nil {
		return ""
	}
	if dto.Nickname != "" {
		return dto.Nickname
	}
	return dto.Username
}

func NewAppContext() (*AppContext, error) {
	gormDB, err := db.NewDB()
	if err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
		return nil, err
	}

	redisStore := datastore.NewRedisClient()

	userRepo := repository.NewUserRepo(gormDB)
	if err := userRepo.Migrate(); err != nil {
		logger.Errorf("Failed to migrate user repository: %v", err)
		return nil, err
	}

	rbacRepo := repository.NewRBACRepo(gormDB)

	if err := seedInitialData(gormDB); err != nil {
		logger.Warn("failed to seed initial data:", err)
	}

	userSrv := service.NewUserServiceWithRBAC(userRepo, rbacRepo, redisStore)
	rbacSrv := service.NewRBACService(rbacRepo)

	// NicknameProvider: 通过 UserService 获取昵称（部署时可通过 gRPC 客户端替换）
	np := &nicknameProvider{userSvc: userSrv}

	// 医疗档案
	medRepo := medicalRepo.NewMedicalRecordRepo(gormDB)
	if err := medRepo.Migrate(); err != nil {
		logger.Errorf("Failed to migrate medical repository: %v", err)
		return nil, err
	}
	medSrv := medicalSvc.NewMedicalRecordService(medRepo, np, "./uploads")

	return &AppContext{
		DB:                   gormDB,
		Redis:                redisStore,
		UserService:          userSrv,
		RBACService:          rbacSrv,
		MedicalRecordService: medSrv,
	}, nil
}

func seedInitialData(gormDB *gorm.DB) error {
	adminRole := &model.Role{
		ID:          1,
		Name:        "管理员",
		Description: "系统管理员，拥有所有权限",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	userRole := &model.Role{
		ID:          2,
		Name:        "普通用户",
		Description: "普通用户，拥有基本权限",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := gormDB.FirstOrCreate(adminRole, model.Role{ID: 1}).Error; err != nil {
		return err
	}
	if err := gormDB.FirstOrCreate(userRole, model.Role{ID: 2}).Error; err != nil {
		return err
	}

	permissions := []model.Permission{
		{ID: 1, Name: "查看用户", Code: "user:view", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "新增用户", Code: "user:create", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Name: "编辑用户", Code: "user:update", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, Name: "删除用户", Code: "user:delete", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, Name: "管理角色", Code: "role:manage", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, p := range permissions {
		perm := p
		if err := gormDB.FirstOrCreate(&perm, model.Permission{ID: p.ID}).Error; err != nil {
			return err
		}
	}

	allPermIDs := []int64{1, 2, 3, 4, 5}
	for _, pid := range allPermIDs {
		rp := model.RolePermission{RoleID: 1, PermissionID: pid}
		gormDB.FirstOrCreate(&rp)
	}
	rpUser := model.RolePermission{RoleID: 2, PermissionID: 1}
	gormDB.FirstOrCreate(&rpUser)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminUser := &model.User{
		UserID:    100000000000000001,
		Username:  "admin",
		Nickname:  "管理员",
		Password:  string(hashedBytes),
		Status:    model.UserStatusActive,
		IsSystem:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := gormDB.FirstOrCreate(adminUser, model.User{UserID: 100000000000000001}).Error; err != nil {
		return err
	}

	ur := model.UserRole{UserID: adminUser.UserID, RoleID: 1}
	gormDB.FirstOrCreate(&ur)

	return nil
}
