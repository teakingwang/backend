package service

import (
	"context"
	"errors"
	"github.com/teakingwang/backend/internal/user/model"
	"strconv"

	"github.com/teakingwang/backend/internal/user/repository"
	"github.com/teakingwang/backend/pkg/auth"
	"github.com/teakingwang/backend/pkg/crypto"
	"github.com/teakingwang/backend/pkg/datastore"
	"github.com/teakingwang/backend/pkg/logger"
	"github.com/teakingwang/backend/pkg/utils/idgen"
)

type UserDTO struct {
	UserID     int64            `json:"user_id"`
	Username   string           `json:"username"`
	Nickname   string           `json:"nickname"`
	Email      string           `json:"email"`
	Mobile     string           `json:"mobile"`
	Status     model.UserStatus `json:"status"`
	StatusName string           `json:"status_name"`
	RoleName   string           `json:"role_name"`
	RoleID     int64            `json:"role_id"`
	Roles      []*model.Role    `json:"roles,omitempty"`
}

type UserService interface {
	Login(ctx context.Context, username, password string) (string, *UserDTO, error)
	GetUserList(ctx context.Context) ([]*UserDTO, int64, error)
	GetUser(ctx context.Context, id int64) (*UserDTO, error)
	CreateUser(ctx context.Context, username, password, email, nickname string, roleID int64) (*UserDTO, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, id int64, password, email, nickname string, roleID int64, status *model.UserStatus) (*UserDTO, error)
}

type userService struct {
	userRepo repository.UserRepo
	rbacRepo repository.RBACRepo
	redis    datastore.Store
}

func NewUserService(userRepo repository.UserRepo, redis datastore.Store) UserService {
	return &userService{userRepo: userRepo, redis: redis}
}

func NewUserServiceWithRBAC(userRepo repository.UserRepo, rbacRepo repository.RBACRepo, redis datastore.Store) UserService {
	return &userService{userRepo: userRepo, rbacRepo: rbacRepo, redis: redis}
}

func (s *userService) Login(ctx context.Context, username, password string) (string, *UserDTO, error) {
	logger.Info("Login called with username:", username)

	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", nil, err
	}

	if user == nil {
		return "", nil, errors.New("用户不存在")
	}

	b, err := crypto.ComparePassword(user.Password, password)
	if err != nil {
		return "", nil, err
	}

	if !b {
		return "", nil, errors.New("密码错误")
	}

	token, err := auth.GenerateToken(strconv.Itoa(int(user.UserID)), user.Username, user.Nickname)
	if err != nil {
		return "", nil, err
	}

	dto := &UserDTO{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Nickname: user.Nickname,
	}

	if s.rbacRepo != nil {
		roles, _ := s.rbacRepo.GetUserRoles(ctx, user.UserID)
		dto.Roles = roles
		if len(roles) > 0 {
			dto.RoleName = roles[0].Name
			dto.RoleID = roles[0].ID
		}
	}

	return token, dto, nil
}

func (s *userService) GetUserList(ctx context.Context) ([]*UserDTO, int64, error) {
	logger.Info("GetUserList called")
	urs, total, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return []*UserDTO{}, 0, nil
	}

	var dtos []*UserDTO
	for _, ur := range urs {
		dto := &UserDTO{
			UserID:     ur.UserID,
			Username:   ur.Username,
			Nickname:   ur.Nickname,
			Email:      ur.Email,
			Mobile:     ur.Mobile,
			Status:     ur.Status,
			StatusName: ur.Status.ToText(),
		}

		if s.rbacRepo != nil {
			roles, _ := s.rbacRepo.GetUserRoles(ctx, ur.UserID)
			if len(roles) > 0 {
				dto.RoleName = roles[0].Name
				dto.RoleID = roles[0].ID
			} else {
				dto.RoleName = "未分配"
			}
		}

		dtos = append(dtos, dto)
	}

	return dtos, total, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*UserDTO, error) {
	logger.Info("GetUser called with ID:", id)
	ur, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if ur == nil {
		return nil, nil
	}

	dto := &UserDTO{
		UserID:     ur.UserID,
		Username:   ur.Username,
		Email:      ur.Email,
		Nickname:   ur.Nickname,
		Mobile:     ur.Mobile,
		Status:     ur.Status,
		StatusName: ur.Status.ToText(),
	}

	if s.rbacRepo != nil {
		roles, _ := s.rbacRepo.GetUserRoles(ctx, ur.UserID)
		dto.Roles = roles
		if len(roles) > 0 {
			dto.RoleName = roles[0].Name
			dto.RoleID = roles[0].ID
		}
	}

	return dto, nil
}

func (s *userService) CreateUser(ctx context.Context, username, password, email, nickname string, roleID int64) (*UserDTO, error) {
	logger.Info("CreateUser called with username:", username)

	existingUser, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UserID:   idgen.NewID(),
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Nickname: nickname,
		Status:   model.UserStatusActive,
		IsSystem: false,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	dto := &UserDTO{
		UserID:     user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		Nickname:   user.Nickname,
		Mobile:     user.Mobile,
		Status:     user.Status,
		StatusName: user.Status.ToText(),
	}

	if s.rbacRepo != nil && roleID > 0 {
		_ = s.rbacRepo.SetUserRole(ctx, user.UserID, roleID)
		role, _ := s.rbacRepo.GetRoleByID(ctx, roleID)
		if role != nil {
			dto.RoleName = role.Name
			dto.RoleID = role.ID
		}
	}

	return dto, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	logger.Info("DeleteUser called with ID:", id)

	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("用户不存在")
	}

	return s.userRepo.Delete(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, id int64, password, email, nickname string, roleID int64, status *model.UserStatus) (*UserDTO, error) {
	logger.Info("UpdateUser called with ID:", id)

	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("用户不存在")
	}

	if password != "" {
		hashedPassword, err := crypto.HashPassword(password)
		if err != nil {
			return nil, err
		}
		existingUser.Password = hashedPassword
	}
	if email != "" {
		existingUser.Email = email
	}
	if nickname != "" {
		existingUser.Nickname = nickname
	}
	if status != nil {
		existingUser.Status = *status
	}

	err = s.userRepo.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	dto := &UserDTO{
		UserID:     existingUser.UserID,
		Username:   existingUser.Username,
		Email:      existingUser.Email,
		Nickname:   existingUser.Nickname,
		Status:     existingUser.Status,
		StatusName: existingUser.Status.ToText(),
	}

	if s.rbacRepo != nil && roleID > 0 {
		_ = s.rbacRepo.SetUserRole(ctx, existingUser.UserID, roleID)
	}
	if s.rbacRepo != nil {
		roles, _ := s.rbacRepo.GetUserRoles(ctx, existingUser.UserID)
		dto.Roles = roles
		if len(roles) > 0 {
			dto.RoleName = roles[0].Name
			dto.RoleID = roles[0].ID
		}
	}

	return dto, nil
}
