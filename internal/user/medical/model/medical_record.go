package model

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecord struct {
	ID           int64  `gorm:"column:id;primaryKey"`
	ParentID     int64  `gorm:"column:parent_id;default:0;index"`
	UserID       int64  `gorm:"column:user_id;not null;index"`
	RecordType   int8   `gorm:"column:record_type;not null;default:1"`
	Title        string `gorm:"column:title;not null;size:255"`
	Description  string `gorm:"column:description;type:text"`
	HospitalName string `gorm:"column:hospital_name;size:128"`
	VisitDate    string `gorm:"column:visit_date;size:32"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (MedicalRecord) TableName() string {
	return "tbl_medical_record"
}

const (
	RecordTypeMedical int8 = 1
	RecordTypeInvoice int8 = 2
)

var RecordTypeNames = map[int8]string{
	RecordTypeMedical: "就诊病历",
	RecordTypeInvoice: "电子发票",
}

func (m MedicalRecord) TypeName() string {
	if name, ok := RecordTypeNames[m.RecordType]; ok {
		return name
	}
	return "未知"
}
