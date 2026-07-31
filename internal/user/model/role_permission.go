package model

type RolePermission struct {
	RoleID       int64 `gorm:"column:role_id;primaryKey"`
	PermissionID int64 `gorm:"column:permission_id;primaryKey"`
}

func (RolePermission) TableName() string {
	return "tbl_role_permission"
}
