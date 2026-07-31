package db

import (
	medicalModel "github.com/teakingwang/backend/internal/user/medical/model"
	"github.com/te
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
		&model.UserRole{},
		&medicalModel.MedicalRecord{},
	)
	if err != nil {
		panic(err)
	}
}
