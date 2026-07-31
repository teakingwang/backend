package model

type UserRole struct {
	UserID int64 `gorm:"column:user_id;primaryKey"`
	RoleID int64 `gorm:"column:role_id;primaryKey"`
}

func (UserRole) TableName() string {
	return "tbl_user_role"
}
