package model

import "time"

type MedicalAttachment struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	RecordID  int64  `gorm:"column:record_id;not null;index"`
	FileName  string `gorm:"column:file_name;size:255"`
	FileURL   string `gorm:"column:file_url;size:512"`
	FileSize  int64  `gorm:"column:file_size;default:0"`
	CreatedAt time.Time
}

func (MedicalAttachment) TableName() string {
	return "tbl_medical_attachment"
}