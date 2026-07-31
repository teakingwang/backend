package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/teakingwang/backend/internal/medical/model"
	"github.com/teakingwang/backend/internal/medical/repository"
	"github.com/teakingwang/backend/pkg/logger"
	"github.com/teakingwang/backend/pkg/utils/idgen"
)

type MedicalRecordDTO struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	RecordType     int8   `json:"record_type"`
	RecordTypeName string `json:"record_type_name"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	HospitalName   string `json:"hospital_name"`
	VisitDate      string `json:"visit_date"`
	FileName       string `json:"file_name"`
	FileURL        string `json:"file_url"`
	FileSize       int64  `json:"file_size"`
	CreatedAt      string `json:"created_at"`
}

type MedicalRecordService interface {
	GetByID(ctx context.Context, id int64) (*MedicalRecordDTO, error)
	ListByUserID(ctx context.Context, userID int64) ([]*MedicalRecordDTO, error)
	Create(ctx context.Context, userID int64, recordType int8, title, description, hospitalName, visitDate string, file *multipart.FileHeader) (*MedicalRecordDTO, error)
	Delete(ctx context.Context, id int64) error
}

type medicalRecordService struct {
	repo      repository.MedicalRecordRepo
	uploadDir string
}

func NewMedicalRecordService(repo repository.MedicalRecordRepo, uploadDir string) MedicalRecordService {
	return &medicalRecordService{repo: repo, uploadDir: uploadDir}
}

func (s *medicalRecordService) GetByID(ctx context.Context, id int64) (*MedicalRecordDTO, error) {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, errors.New("档案不存在")
	}
	return toDTO(record), nil
}

func (s *medicalRecordService) ListByUserID(ctx context.Context, userID int64) ([]*MedicalRecordDTO, error) {
	records, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	dtos := make([]*MedicalRecordDTO, 0, len(records))
	for _, r := range records {
		dtos = append(dtos, toDTO(r))
	}
	return dtos, nil
}

func (s *medicalRecordService) Create(ctx context.Context, userID int64, recordType int8, title, description, hospitalName, visitDate string, file *multipart.FileHeader) (*MedicalRecordDTO, error) {
	recordID := idgen.NewID()

	var fileName, fileURL string
	var fileSize int64

	if file != nil {
		ext := filepath.Ext(file.Filename)
		fileName = fmt.Sprintf("%d_%s%s", recordID, uuid.New().String()[:8], ext)
		fileURL = "/uploads/" + fileName
		fileSize = file.Size

		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("打开上传文件失败: %w", err)
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join(s.uploadDir, fileName))
		if err != nil {
			return nil, fmt.Errorf("创建文件失败: %w", err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return nil, fmt.Errorf("保存文件失败: %w", err)
		}
	}

	record := &model.MedicalRecord{
		ID:           recordID,
		UserID:       userID,
		RecordType:   recordType,
		Title:        title,
		Description:  description,
		HospitalName: hospitalName,
		VisitDate:    visitDate,
		FileName:     fileName,
		FileURL:      fileURL,
		FileSize:     fileSize,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, record); err != nil {
		logger.Errorf("create medical record failed: %v", err)
		return nil, err
	}

	return toDTO(record), nil
}

func (s *medicalRecordService) Delete(ctx context.Context, id int64) error {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("档案不存在")
	}
	return s.repo.Delete(ctx, id)
}

func toDTO(r *model.MedicalRecord) *MedicalRecordDTO {
	return &MedicalRecordDTO{
		ID:             r.ID,
		UserID:         r.UserID,
		RecordType:     r.RecordType,
		RecordTypeName: r.TypeName(),
		Title:          r.Title,
		Description:    r.Description,
		HospitalName:   r.HospitalName,
		VisitDate:      r.VisitDate,
		FileName:       r.FileName,
		FileURL:        r.FileURL,
		FileSize:       r.FileSize,
		CreatedAt:      r.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
