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

type NicknameProvider interface {
	GetUserNickname(ctx context.Context, userID int64) string
}

type AttachmentDTO struct {
	ID       int64  `json:"id"`
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	FileSize int64  `json:"file_size"`
}

type MedicalRecordDTO struct {
	ID             int64               `json:"id"`
	ParentID       int64               `json:"parent_id"`
	UserID         int64               `json:"user_id"`
	UserNickname   string              `json:"user_nickname"`
	RecordType     int8                `json:"record_type"`
	RecordTypeName string              `json:"record_type_name"`
	Title          string              `json:"title"`
	Description    string              `json:"description"`
	HospitalName   string              `json:"hospital_name"`
	VisitDate      string              `json:"visit_date"`
	Attachments    []*AttachmentDTO    `json:"attachments"`
	Invoices       []*MedicalRecordDTO `json:"invoices,omitempty"`
	CreatedAt      string              `json:"created_at"`
}

type MedicalRecordService interface {
	GetByID(ctx context.Context, id int64) (*MedicalRecordDTO, error)
	ListByUserID(ctx context.Context, userID int64, isAdmin bool) ([]*MedicalRecordDTO, error)
	Create(ctx context.Context, userID int64, recordType int8, title, description, hospitalName, visitDate string, files []*multipart.FileHeader, invoiceFiles []*multipart.FileHeader) (*MedicalRecordDTO, error)
	Update(ctx context.Context, id int64, title, description, hospitalName, visitDate string, recordType int8, files []*multipart.FileHeader, invoiceFiles []*multipart.FileHeader) (*MedicalRecordDTO, error)
	Delete(ctx context.Context, id int64) error
}

type medicalRecordService struct {
	repo             repository.MedicalRecordRepo
	nicknameProvider NicknameProvider
	uploadDir        string
}

func NewMedicalRecordService(repo repository.MedicalRecordRepo, nicknameProvider NicknameProvider, uploadDir string) MedicalRecordService {
	return &medicalRecordService{repo: repo, nicknameProvider: nicknameProvider, uploadDir: uploadDir}
}

func (s *medicalRecordService) populateNicknames(ctx context.Context, dtos []*MedicalRecordDTO) {
	if s.nicknameProvider == nil {
		return
	}
	for _, d := range dtos {
		d.UserNickname = s.nicknameProvider.GetUserNickname(ctx, d.UserID)
	}
}

func (s *medicalRecordService) saveFiles(files []*multipart.FileHeader, prefix int64) ([]*model.MedicalAttachment, error) {
	var attachments []*model.MedicalAttachment
	for _, f := range files {
		ext := filepath.Ext(f.Filename)
		name := fmt.Sprintf("%d_%s%s", idgen.NewID(), uuid.New().String()[:8], ext)
		url := "/uploads/" + name

		src, err := f.Open()
		if err != nil {
			return nil, err
		}
		dst, err := os.Create(filepath.Join(s.uploadDir, name))
		if err != nil {
			src.Close()
			return nil, err
		}
		io.Copy(dst, src)
		src.Close()
		dst.Close()

		attachments = append(attachments, &model.MedicalAttachment{
			ID:       idgen.NewID(),
			RecordID: prefix,
			FileName: f.Filename,
			FileURL:  url,
			FileSize: f.Size,
		})
	}
	return attachments, nil
}

func (s *medicalRecordService) saveInvoiceChildRecords(ctx context.Context, parentID int64, userID int64, invoiceFiles []*multipart.FileHeader) ([]*MedicalRecordDTO, error) {
	var invoiceDTOs []*MedicalRecordDTO
	for _, f := range invoiceFiles {
		invoiceID := idgen.NewID()

		atts, err := s.saveFiles([]*multipart.FileHeader{f}, invoiceID)
		if err != nil {
			logger.Errorf("save invoice file failed: %v", err)
			continue
		}

		record := &model.MedicalRecord{
			ID:           invoiceID,
			ParentID:     parentID,
			UserID:       userID,
			RecordType:   model.RecordTypeInvoice,
			Title:        fmt.Sprintf("发票-%s", f.Filename),
			HospitalName: "",
			VisitDate:    "",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if err := s.repo.Create(ctx, record); err != nil {
			logger.Errorf("create invoice record failed: %v", err)
			continue
		}
		if len(atts) > 0 {
			for i := range atts {
				atts[i].RecordID = invoiceID
			}
			if err := s.repo.CreateAttachments(ctx, atts); err != nil {
				logger.Errorf("create invoice attachments failed: %v", err)
			}
		}
		invoiceDTOs = append(invoiceDTOs, toDTO(record, atts))
	}
	return invoiceDTOs, nil
}

func (s *medicalRecordService) GetByID(ctx context.Context, id int64) (*MedicalRecordDTO, error) {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil || record == nil {
		if record == nil {
			return nil, errors.New("档案不存在")
		}
		return nil, err
	}
	atts, _ := s.repo.GetAttachments(ctx, id)
	dto := toDTO(record, atts)

	children, _ := s.repo.ListByParentID(ctx, id)
	if len(children) > 0 {
		invoices := make([]*MedicalRecordDTO, 0, len(children))
		for _, child := range children {
			childAtts, _ := s.repo.GetAttachments(ctx, child.ID)
			invoices = append(invoices, toDTO(child, childAtts))
		}
		dto.Invoices = invoices
	}

	s.populateNicknames(ctx, []*MedicalRecordDTO{dto})
	return dto, nil
}

func (s *medicalRecordService) ListByUserID(ctx context.Context, userID int64, isAdmin bool) ([]*MedicalRecordDTO, error) {
	var records []*model.MedicalRecord
	var err error
	if isAdmin {
		records, err = s.repo.ListAll(ctx)
	} else {
		records, err = s.repo.ListByUserID(ctx, userID)
	}
	if err != nil {
		return nil, err
	}
	dtos := make([]*MedicalRecordDTO, 0, len(records))
	for _, r := range records {
		atts, _ := s.repo.GetAttachments(ctx, r.ID)
		dto := toDTO(r, atts)
		children, _ := s.repo.ListByParentID(ctx, r.ID)
		if len(children) > 0 {
			invoices := make([]*MedicalRecordDTO, 0, len(children))
			for _, child := range children {
				childAtts, _ := s.repo.GetAttachments(ctx, child.ID)
				invoices = append(invoices, toDTO(child, childAtts))
			}
			dto.Invoices = invoices
		}
		dtos = append(dtos, dto)
	}

	s.populateNicknames(ctx, dtos)
	return dtos, nil
}

func (s *medicalRecordService) Create(ctx context.Context, userID int64, recordType int8, title, description, hospitalName, visitDate string, files []*multipart.FileHeader, invoiceFiles []*multipart.FileHeader) (*MedicalRecordDTO, error) {
	recordID := idgen.NewID()
	record := &model.MedicalRecord{
		ID:           recordID,
		ParentID:     0,
		UserID:       userID,
		RecordType:   recordType,
		Title:        title,
		Description:  description,
		HospitalName: hospitalName,
		VisitDate:    visitDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.repo.Create(ctx, record); err != nil {
		return nil, err
	}

	var atts []*model.MedicalAttachment
	if len(files) > 0 {
		var err error
		atts, err = s.saveFiles(files, recordID)
		if err != nil {
			logger.Errorf("save files failed: %v", err)
		} else {
			for i := range atts {
				atts[i].RecordID = recordID
			}
			if err := s.repo.CreateAttachments(ctx, atts); err != nil {
				logger.Errorf("create attachments failed: %v", err)
			}
		}
	}

	dto := toDTO(record, atts)

	if len(invoiceFiles) > 0 {
		invoices, err := s.saveInvoiceChildRecords(ctx, recordID, userID, invoiceFiles)
		if err != nil {
			logger.Errorf("save invoice child records failed: %v", err)
		}
		dto.Invoices = invoices
	}

	return dto, nil
}

func (s *medicalRecordService) Update(ctx context.Context, id int64, title, description, hospitalName, visitDate string, recordType int8, files []*multipart.FileHeader, invoiceFiles []*multipart.FileHeader) (*MedicalRecordDTO, error) {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil || record == nil {
		if record == nil {
			return nil, errors.New("档案不存在")
		}
		return nil, err
	}

	if title != "" {
		record.Title = title
	}
	if description != "" {
		record.Description = description
	}
	if hospitalName != "" {
		record.HospitalName = hospitalName
	}
	if visitDate != "" {
		record.VisitDate = visitDate
	}
	if recordType > 0 {
		record.RecordType = recordType
	}
	record.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, record); err != nil {
		return nil, err
	}

	if len(files) > 0 {
		s.repo.DeleteAttachments(ctx, id)
		atts, err := s.saveFiles(files, id)
		if err != nil {
			logger.Errorf("save files failed: %v", err)
		} else {
			s.repo.CreateAttachments(ctx, atts)
		}
	}

	if len(invoiceFiles) > 0 {
		s.saveInvoiceChildRecords(ctx, id, record.UserID, invoiceFiles)
	}

	atts, _ := s.repo.GetAttachments(ctx, id)
	dto := toDTO(record, atts)

	children, _ := s.repo.ListByParentID(ctx, id)
	if len(children) > 0 {
		invoices := make([]*MedicalRecordDTO, 0, len(children))
		for _, child := range children {
			childAtts, _ := s.repo.GetAttachments(ctx, child.ID)
			invoices = append(invoices, toDTO(child, childAtts))
		}
		dto.Invoices = invoices
	}

	return dto, nil
}

func (s *medicalRecordService) Delete(ctx context.Context, id int64) error {
	record, err := s.repo.GetByID(ctx, id)
	if err != nil || record == nil {
		if record == nil {
			return errors.New("档案不存在")
		}
		return err
	}

	children, _ := s.repo.ListByParentID(ctx, id)
	for _, child := range children {
		s.repo.DeleteAttachments(ctx, child.ID)
	}
	s.repo.DeleteByParentID(ctx, id)

	s.repo.DeleteAttachments(ctx, id)
	return s.repo.Delete(ctx, id)
}

func toDTO(r *model.MedicalRecord, atts []*model.MedicalAttachment) *MedicalRecordDTO {
	aDTOs := make([]*AttachmentDTO, 0, len(atts))
	for _, a := range atts {
		aDTOs = append(aDTOs, &AttachmentDTO{
			ID:       a.ID,
			FileName: a.FileName,
			FileURL:  a.FileURL,
			FileSize: a.FileSize,
		})
	}
	return &MedicalRecordDTO{
		ID:             r.ID,
		ParentID:       r.ParentID,
		UserID:         r.UserID,
		RecordType:     r.RecordType,
		RecordTypeName: r.TypeName(),
		Title:          r.Title,
		Description:    r.Description,
		HospitalName:   r.HospitalName,
		VisitDate:      r.VisitDate,
		Attachments:    aDTOs,
		CreatedAt:      r.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}