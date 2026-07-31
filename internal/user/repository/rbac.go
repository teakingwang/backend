package repository

import (
	"context"

	"github.com/teakingwang/backend/internal/user/model"
	"gorm.io/gorm"
)

type RBACRepo interface {
	// 角色
	GetAllRoles(ctx context.Context) ([]*model.Role, error)
	GetRoleByID(ctx context.Context, id int64) (*model.Role, error)
	CreateRole(ctx context.Context, role *model.Role) error
	UpdateRole(ctx context.Context, role *model.Role) error
	DeleteRole(ctx context.Context, id int64) error

	// 权限
	GetAllPermissions(ctx context.Context) ([]*model.Permission, error)
	GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]*model.Permission, error)

	// 角色-权限关联
	SetRolePermissions(ctx context.Context, roleID int64, permIDs []int64) error

	// 用户-角色关联
	GetUserRoles(ctx context.Context, userID int64) ([]*model.Role, error)
	SetUserRole(ctx context.Context, userID int64, roleID int64) error
	RemoveUserRole(ctx context.Context, userID int64, roleID int64) error

	// 权限检查
	CheckUserPermission(ctx context.Context, userID int64, permCode string) (bool, error)
}

type rbacRepo struct {
	db *gorm.DB
}

func NewRBACRepo(gormDB *gorm.DB) RBACRepo {
	return &rbacRepo{db: gormDB}
}

func (r *rbacRepo) GetAllRoles(ctx context.Context) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).Find(&roles).Error
	return roles, err
}

func (r *rbacRepo) GetRoleByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *rbacRepo) CreateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *rbacRepo) UpdateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *rbacRepo) DeleteRole(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
}

func (r *rbacRepo) GetAllPermissions(ctx context.Context) ([]*model.Permission, error) {
	var perms []*model.Permission
	err := r.db.WithContext(ctx).Find(&perms).Error
	return perms, err
}

func (r *rbacRepo) GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]*model.Permission, error) {
	var perms []*model.Permission
	err := r.db.WithContext(ctx).
		Joins("JOIN tbl_role_permission rp ON rp.permission_id = tbl_permission.id").
		Where("rp.role_id = ?", roleID).
		Find(&perms).Error
	return perms, err
}

func (r *rbacRepo) SetRolePermissions(ctx context.Context, roleID int64, permIDs []int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}
		for _, pid := range permIDs {
			rp := model.RolePermission{RoleID: roleID, PermissionID: pid}
			if err := tx.Create(&rp).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *rbacRepo) GetUserRoles(ctx context.Context, userID int64) ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.WithContext(ctx).
		Joins("JOIN tbl_user_role ur ON ur.role_id = tbl_role.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *rbacRepo) SetUserRole(ctx context.Context, userID int64, roleID int64) error {
	ur := model.UserRole{UserID: userID, RoleID: roleID}
	return r.db.WithContext(ctx).Create(&ur).Error
}

func (r *rbacRepo) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&model.UserRole{}).Error
}

func (r *rbacRepo) CheckUserPermission(ctx context.Context, userID int64, permCode string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Table("tbl_user_role ur").
		Joins("JOIN tbl_role_permission rp ON rp.role_id = ur.role_id").
		Joins("JOIN tbl_permission p ON p.id = rp.permission_id").
		Where("ur.user_id = ? AND p.code = ?", userID, permCode).
		Count(&count).Error
	return count > 0, err
}
