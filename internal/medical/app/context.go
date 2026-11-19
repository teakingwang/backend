package app

import (
	"github.com/teakingwang/backend/config"
	"github.com/teakingwang/backend/internal/medical/repository"
	"github.com/teakingwang/backend/internal/medical/service"
	"github.com/teakingwang/backend/internal/medical/userclient"
	"github.com/teakingwang/backend/pkg/consul"
	"github.com/teakingwang/backend/pkg/db"
	"github.com/teakingwang/backend/pkg/logger"
	"gorm.io/gorm"
)

type AppContext struct {
	DB                   *gorm.DB
	MedicalRecordService service.MedicalRecordService
	UserClient           *userclient.UserClient
}

func NewAppContext(uploadDir string) (*AppContext, error) {
	gormDB, err := db.NewDB()
	if err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
		return nil, err
	}

	medicalRepo := repository.NewMedicalRecordRepo(gormDB)
	if err := medicalRepo.Migrate(); err != nil {
		logger.Errorf("Failed to migrate medical repository: %v", err)
		return nil, err
	}

	var nicknameProvider service.NicknameProvider
	consulDiscovery, discoveryErr := consul.NewConsulDiscovery(config.Config.Consul.Address)
	if discoveryErr != nil {
		logger.Warnf("Failed to create consul discovery: %v", discoveryErr)
	} else {
		uc, err := userclient.NewUserClient(consulDiscovery)
		if err != nil {
			logger.Warnf("Failed to create user client: %v", err)
		} else {
			nicknameProvider = uc
		}
	}

	medicalSvc := service.NewMedicalRecordService(medicalRepo, nicknameProvider, uploadDir)

	return &AppContext{
		DB:                   gormDB,
		MedicalRecordService: medicalSvc,
	}, nil
}