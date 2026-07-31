package service

import (
	"context"

	"github.com/teakingwang/backend/internal/user/model"
	"github.com/teakingwang/backend/internal/user/repository"
)

type RBACService interface {
	// 角色管理
	GetAllRoles(ctx context.Context) ([]*model.Role, error)
	GetRoleByID(ctx context.Context, id int64) (*model.Role, error)
	CreateRole(ctx context.Context, name, description string) (*model.Role, error)
	UpdateRole(ctx context.Context, id int64, name, description string) (*model.Role, error)
	DeleteRole(ctx context.Context, id int64) error

	// 权限管理
	GetAllPermissions(ctx context.Context) ([]*model.Permission, error)
	GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]*model.Permission, error)
	SetRolePermissions(ctx context.Context, roleID int64, permIDs []int64) error

	// 用户角色
	GetUserRoles(ctx context.Context, userID int64) ([]*model.Role, error)
	SetUserRole(ctx context.Context, userID int64, roleID int64) error

	// 权限检查
	CheckPermission(ctx context.Context, userID int64, permCode string) (bool, error)
}

type rbacService struct {
	rbacRepo repository.RBACRepo
}

func NewRBACService(rbacRepo repository.RBACRepo) RBACService {
	return &rbacService{rbacRepo: rbacRepo}
}

func (s *rbacService) GetAllRoles(ctx context.Context) ([]*model.Role, error) {
	return s.rbacRepo.GetAllRoles(ctx)
}

func (s *rbacService) GetRoleByID(ctx context.Context, id int64) (*model.Role, error) {
	return s.rbacRepo.GetRoleByID(ctx, id)
}

func (s *rbacService) CreateRole(ctx context.Context, name, description string) (*model.Role, error) {
	role := &model.Role{
		Name:        name,
		Description: description,
	}
	err := s.rbacRepo.CreateRole(ctx, role)
	return role, err
}

func (s *rbacService) UpdateRole(ctx context.Context, id int64, name, description string) (*model.Role, error) {
	role, err := s.rbacRepo.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		role.Name = name
	}
	if description != "" {
		role.Description = description
	}
	err = s.rbacRepo.UpdateRole(ctx, role)
	return role, err
}

func (s *rbacService) DeleteRole(ctx context.Context, id int64) error {
	return s.rbacRepo.DeleteRole(ctx, id)
}

func (s *rbacService) GetAllPermissions(ctx context.Context) ([]*model.Permission, error) {
	return s.rbacRepo.GetAllPermissions(ctx)
}

func (s *rbacService) GetPermissionsByRoleID(ctx context.Context, roleID int64) ([]*model.Permission, error) {
	return s.rbacRepo.GetPermissionsByRoleID(ctx, roleID)
}

func (s *rbacService) SetRolePermissions(ctx context.Context, roleID int64, permIDs []int64) error {
	return s.rbacRepo.SetRolePermissions(ctx, roleID, permIDs)
}

func (s *rbacService) GetUserRoles(ctx context.Context, userID int64) ([]*model.Role, error) {
	return s.rbacRepo.GetUserRoles(ctx, userID)
}

func (s *rbacService) SetUserRole(ctx context.Context, userID int64, roleID int64) error {
	return s.rbacRepo.SetUserRole(ctx, userID, roleID)
}

func (s *rbacService) CheckPermission(ctx context.Context, userID int64, permCode string) (bool, error) {
	return s.rbacRepo.CheckUserPermission(ctx, userID, permCode)
}
