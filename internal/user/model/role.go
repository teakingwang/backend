package model

import (
	"time"
)

type Role struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:name;not null;uniqueIndex;size:64"`
	Description string `gorm:"column:description;size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Role) TableName() string {
	return "tbl_role"
}
