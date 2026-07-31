package model

import (
	"time"
)

type Permission struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;not null;size:64"`
	Code      string `gorm:"column:code;not null;uniqueIndex;size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Permission) TableName() string {
	return "tbl_permission"
}
