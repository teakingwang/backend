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
	ListAll(ctx context.Context) ([]*model.MedicalRecord, error)
	ListByParentID(ctx context.Context, parentID int64) ([]*model.MedicalRecord, error)
	Create(ctx context.Context, record *model.MedicalRecord) error
	Update(ctx context.Context, record *model.MedicalRecord) error
	Delete(ctx context.Context, id int64) error
	DeleteByParentID(ctx context.Context, parentID int64) error
	GetAttachments(ctx context.Context, recordID int64) ([]*model.MedicalAttachment, error)
	CreateAttachments(ctx context.Context, attachments []*model.MedicalAttachment) error
	DeleteAttachments(ctx context.Context, recordID int64) error
}

type medicalRecordRepo struct {
	db *gorm.DB
}

func NewMedicalRecordRepo(gormDB *gorm.DB) MedicalRecordRepo {
	return &medicalRecordRepo{db: gormDB}
}

func (r *medicalRecordRepo) Migrate() error {
	return r.db.AutoMigrate(&model.MedicalRecord{}, &model.MedicalAttachment{})
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
		Where("user_id = ? AND parent_id = 0", userID).
		Order("visit_date DESC, created_at DESC").
		Find(&records).Error
	return records, err
}

func (r *medicalRecordRepo) ListAll(ctx context.Context) ([]*model.MedicalRecord, error) {
	var records []*model.MedicalRecord
	err := r.db.WithContext(ctx).
		Where("parent_id = 0").
		Order("visit_date DESC, created_at DESC").
		Find(&records).Error
	return records, err
}

func (r *medicalRecordRepo) ListByParentID(ctx context.Context, parentID int64) ([]*model.MedicalRecord, error) {
	var records []*model.MedicalRecord
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("created_at ASC").
		Find(&records).Error
	return records, err
}

func (r *medicalRecordRepo) Create(ctx context.Context, record *model.MedicalRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *medicalRecordRepo) Update(ctx context.Context, record *model.MedicalRecord) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *medicalRecordRepo) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MedicalRecord{}).Error
}

func (r *medicalRecordRepo) DeleteByParentID(ctx context.Context, parentID int64) error {
	return r.db.WithContext(ctx).Where("parent_id = ?", parentID).Delete(&model.MedicalRecord{}).Error
}

func (r *medicalRecordRepo) GetAttachments(ctx context.Context, recordID int64) ([]*model.MedicalAttachment, error) {
	var atts []*model.MedicalAttachment
	err := r.db.WithContext(ctx).Where("record_id = ?", recordID).Find(&atts).Error
	return atts, err
}

func (r *medicalRecordRepo) CreateAttachments(ctx context.Context, attachments []*model.MedicalAttachment) error {
	return r.db.WithContext(ctx).Create(attachments).Error
}

func (r *medicalRecordRepo) DeleteAttachments(ctx context.Context, recordID int64) error {
	return r.db.WithContext(ctx).Where("record_id = ?", recordID).Delete(&model.MedicalAttachment{}).Error
}