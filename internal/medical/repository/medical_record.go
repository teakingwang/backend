package repository

import (
	"context"
	"errors"

	"github.com/teakingwang/backend/internal/medical/model"
	"gorm.io/gorm"
)

type MedicalRecordRepo interface {
	Migrate() error
	GetByID(ctx context.Context, id int64) (*model.MedicalRecord, error)
	ListByUserID(ctx context.Context, userID int64) ([]*model.MedicalRecord, error)
	Create(ctx context.Context, record *model.MedicalRecord) error
	Delete(ctx context.Context, id int64) error
}

type medicalRecordRepo struct {
	db *gorm.DB
}

func NewMedicalRecordRepo(gormDB *gorm.DB) MedicalRecordRepo {
	return &medicalRecordRepo{db: gormDB}
}

func (r *medicalRecordRepo) Migrate() error {
	return r.db.AutoMigrate(&model.MedicalRecord{})
}

func (r *medicalRecordRepo) GetByID(ctx context.Context, id int64) (*model.MedicalRecord, error) {
	var m model.MedicalRecord
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &m, err
}

func (r *medicalRecordRepo) ListByUserID(ctx context.Context, userID int64) ([]*model.MedicalRecord, error) {
	var records []*model.MedicalRecord
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&records).Error
	return records, err
}

func (r *medicalRecordRepo) Create(ctx context.Context, record *model.MedicalRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *medicalRecordRepo) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MedicalRecord{}).Error
}
