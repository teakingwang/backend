package handler

import (
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/user/medical/service"
	"github.com/teakingwang/backend/pkg/errs"
	"github.com/teakingwang/backend/pkg/logger"
)

type MedicalHandler struct {
	svc service.MedicalRecordService
}

func NewMedicalHandler(svc service.MedicalRecordService) *MedicalHandler {
	return &MedicalHandler{svc: svc}
}

type AttachmentItem struct {
	ID       int64  `json:"id"`
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	FileSize int64  `json:"file_size"`
}

type MedicalRecordItem struct {
	ID             string               `json:"id"`
	ParentID       string               `json:"parent_id"`
	UserID         string               `json:"user_id"`
	UserNickname   string               `json:"user_nickname"`
	RecordType     int8                 `json:"record_type"`
	RecordTypeName string               `json:"record_type_name"`
	Title          string               `json:"title"`
	Description    string               `json:"description"`
	HospitalName   string               `json:"hospital_name"`
	VisitDate      string               `json:"visit_date"`
	Attachments    []*AttachmentItem    `json:"attachments"`
	Invoices       []*MedicalRecordItem `json:"invoices,omitempty"`
	CreatedAt      string               `json:"created_at"`
}

func dtoToItem(dto *service.MedicalRecordDTO) *MedicalRecordItem {
	if dto == nil {
		return nil
	}
	atts := make([]*AttachmentItem, 0, len(dto.Attachments))
	for _, a := range dto.Attachments {
		atts = append(atts, &AttachmentItem{
			ID:       a.ID,
			FileName: a.FileName,
			FileURL:  a.FileURL,
			FileSize: a.FileSize,
		})
	}
	item := &MedicalRecordItem{
		ID:             fmt.Sprintf("%d", dto.ID),
		ParentID:       fmt.Sprintf("%d", dto.ParentID),
		UserID:         fmt.Sprintf("%d", dto.UserID),
		UserNickname:   dto.UserNickname,
		RecordType:     dto.RecordType,
		RecordTypeName: dto.RecordTypeName,
		Title:          dto.Title,
		Description:    dto.Description,
		HospitalName:   dto.HospitalName,
		VisitDate:      dto.VisitDate,
		Attachments:    atts,
		CreatedAt:      dto.CreatedAt,
	}
	if len(dto.Invoices) > 0 {
		invoices := make([]*MedicalRecordItem, 0, len(dto.Invoices))
		for _, inv := range dto.Invoices {
			invoices = append(invoices, dtoToItem(inv))
		}
		item.Invoices = invoices
	}
	return item
}

func getUserID(c *gin.Context) int64 {
	uid, _ := c.Get("user_id")
	uidStr, _ := uid.(string)
	id, _ := strconv.ParseInt(uidStr, 10, 64)
	return id
}

func getFilesFromForm(c *gin.Context, fieldName string) []*multipart.FileHeader {
	form, _ := c.MultipartForm()
	if form != nil {
		files := form.File[fieldName]
		if len(files) > 0 {
			return files
		}
	}
	return nil
}

func (h *MedicalHandler) GetRecord(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	dto, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}
	errs.ResponseSuccessWithData(c, dtoToItem(dto))
}

func (h *MedicalHandler) ListRecords(c *gin.Context) {
	userID := getUserID(c)
	if userID <= 0 {
		errs.ResponseError(c, errs.CodeUnauthorized, "invalid user")
		return
	}
	admin := c.Query("admin") == "1"
	dtos, err := h.svc.ListByUserID(c.Request.Context(), userID, admin)
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
	userID := getUserID(c)
	rt, _ := strconv.Atoi(c.PostForm("record_type"))
	if rt != 1 && rt != 2 {
		rt = 1
	}
	title := c.PostForm("title")
	if title == "" {
		errs.ResponseError(c, errs.CodeInvalidArgs, "title is required")
		return
	}
	description := c.PostForm("description")
	hospitalName := c.PostForm("hospital_name")
	visitDate := c.PostForm("visit_date")

	files := getFilesFromForm(c, "files")
	if len(files) == 0 {
		f, _ := c.FormFile("file")
		if f != nil {
			files = []*multipart.FileHeader{f}
		}
	}
	invoiceFiles := getFilesFromForm(c, "invoice_files")

	dto, err := h.svc.Create(c.Request.Context(), userID, int8(rt), title, description, hospitalName, visitDate, files, invoiceFiles)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}
	errs.ResponseSuccessWithData(c, dtoToItem(dto))
}

func (h *MedicalHandler) UpdateRecord(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	rt, _ := strconv.Atoi(c.PostForm("record_type"))
	title := c.PostForm("title")
	description := c.PostForm("description")
	hospitalName := c.PostForm("hospital_name")
	visitDate := c.PostForm("visit_date")

	files := getFilesFromForm(c, "files")
	if len(files) == 0 {
		f, _ := c.FormFile("file")
		if f != nil {
			files = []*multipart.FileHeader{f}
		}
	}
	invoiceFiles := getFilesFromForm(c, "invoice_files")

	dto, err := h.svc.Update(c.Request.Context(), id, title, description, hospitalName, visitDate, int8(rt), files, invoiceFiles)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}
	errs.ResponseSuccessWithData(c, dtoToItem(dto))
}

func (h *MedicalHandler) DeleteInvoice(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}
	errs.ResponseSuccess(c)
}

func (h *MedicalHandler) DeleteRecord(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}
	errs.ResponseSuccess(c)
	logger.Info("DeleteRecord done")
}
