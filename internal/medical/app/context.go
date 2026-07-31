package app

import (
	"github.com/teakingwang/backend/internal/medical/repository"
	"github.com/teakingwang/backend/internal/medical/service"
	"github.com/teakingwang/backend/pkg/db"
	"github.com/teakingwang/backend/pkg/logger"
	"gorm.io/gorm"
)

type AppContext struct {
	DB                   *gorm.DB
	MedicalRecordService service.MedicalRecordService
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

	medicalSvc := service.NewMedicalRecordService(medicalRepo, uploadDir)

	return &AppContext{
		DB:                   gormDB,
		MedicalRecordService: medicalSvc,
	}, nil
}
