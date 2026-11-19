package controller

import (
	"context"

	pb "github.com/teakingwang/backend/api/medical"
	"github.com/teakingwang/backend/internal/medical/service"
)

type MedicalController struct {
	svc service.MedicalRecordService
}

func NewMedicalController(svc service.MedicalRecordService) *MedicalController {
	return &MedicalController{svc: svc}
}

func (c *MedicalController) GetMedicalRecord(ctx context.Context, req *pb.GetMedicalRecordRequest) (*pb.GetMedicalRecordResponse, error) {
	dto, err := c.svc.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetMedicalRecordResponse{
		Id:     dto.ID,
		UserId: dto.UserID,
		Title:  dto.Title,
	}, nil
}

func (c *MedicalController) ListMedicalRecords(ctx context.Context, req *pb.ListMedicalRecordsRequest) (*pb.ListMedicalRecordsResponse, error) {
	dtos, err := c.svc.ListByUserID(ctx, req.GetUserId(), false)
	if err != nil {
		return nil, err
	}
	records := make([]*pb.MedicalRecord, 0, len(dtos))
	for _, dto := range dtos {
		records = append(records, &pb.MedicalRecord{
			Id:           dto.ID,
			ParentId:     dto.ParentID,
			UserId:       dto.UserID,
			Title:        dto.Title,
			Description:  dto.Description,
			HospitalName: dto.HospitalName,
			VisitDate:    dto.VisitDate,
		})
	}
	return &pb.ListMedicalRecordsResponse{Records: records}, nil
}

func (c *MedicalController) DeleteMedicalRecord(ctx context.Context, req *pb.DeleteMedicalRecordRequest) (*pb.DeleteMedicalRecordResponse, error) {
	if err := c.svc.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteMedicalRecordResponse{}, nil
}