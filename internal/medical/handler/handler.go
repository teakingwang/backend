package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/medical/service"
	"github.com/teakingwang/backend/pkg/errs"
	"github.com/teakingwang/backend/pkg/logger"
)

type MedicalHandler struct {
	svc service.MedicalRecordService
}

func NewMedicalHandler(svc service.MedicalRecordService) *MedicalHandler {
	return &MedicalHandler{svc: svc}
}

func dtoToItem(dto *service.MedicalRecordDTO) *MedicalRecordItem {
	if dto == nil {
		return nil
	}
	return &MedicalRecordItem{
		ID:             fmt.Sprintf("%d", dto.ID),
		UserID:         fmt.Sprintf("%d", dto.UserID),
		RecordType:     dto.RecordType,
		RecordTypeName: dto.RecordTypeName,
		Title:          dto.Title,
		Description:    dto.Description,
		HospitalName:   dto.HospitalName,
		VisitDate:      dto.VisitDate,
		FileName:       dto.FileName,
		FileURL:        dto.FileURL,
		FileSize:       dto.FileSize,
		CreatedAt:      dto.CreatedAt,
	}
}

func (h *MedicalHandler) GetRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid ID")
		return
	}

	dto, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	errs.ResponseSuccessWithData(c, dtoToItem(dto))
}

func (h *MedicalHandler) ListRecords(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid user_id")
		return
	}

	dtos, err := h.svc.ListByUserID(c.Request.Context(), userID)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	items := make([]*MedicalRecordItem, 0, len(dtos))
	for _, dto := range dtos {
		items = append(items, dtoToItem(dto))
	}

	errs.ResponseSuccessWithData(c, items)
}

func (h *MedicalHandler) CreateRecord(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid user_id")
		return
	}

	recordTypeStr := c.PostForm("record_type")
	recordType, _ := strconv.Atoi(recordTypeStr)
	if recordType != 1 && recordType != 2 {
		recordType = 1
	}

	title := c.PostForm("title")
	if title == "" {
		errs.ResponseError(c, errs.CodeInvalidArgs, "title is required")
		return
	}

	description := c.PostForm("description")
	hospitalName := c.PostForm("hospital_name")
	visitDate := c.PostForm("visit_date")

	file, err := c.FormFile("file")
	if err != nil {
		logger.Warn("no file uploaded:", err)
	}

	dto, err := h.svc.Create(c.Request.Context(), userID, int8(recordType), title, description, hospitalName, visitDate, file)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	errs.ResponseSuccessWithData(c, dtoToItem(dto))
}

func (h *MedicalHandler) DeleteRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid ID")
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	errs.ResponseSuccess(c)
}

type MedicalRecordItem struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
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
