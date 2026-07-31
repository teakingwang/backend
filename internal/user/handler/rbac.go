package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/user/service"
	"github.com/teakingwang/backend/pkg/errs"
)

type RBACHandler struct {
	rbacSvc service.RBACService
}

func NewRBACHandler(rbacSvc service.RBACService) *RBACHandler {
	return &RBACHandler{rbacSvc: rbacSvc}
}

// GetRoles 获取所有角色
func (h *RBACHandler) GetRoles(c *gin.Context) {
	roles, err := h.rbacSvc.GetAllRoles(c.Request.Context())
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	type RoleItem struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	items := make([]*RoleItem, 0, len(roles))
	for _, r := range roles {
		items = append(items, &RoleItem{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
		})
	}

	errs.ResponseSuccessWithData(c, items)
}

// GetPermissions 获取所有权限
func (h *RBACHandler) GetPermissions(c *gin.Context) {
	perms, err := h.rbacSvc.GetAllPermissions(c.Request.Context())
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	type PermItem struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}

	items := make([]*PermItem, 0, len(perms))
	for _, p := range perms {
		items = append(items, &PermItem{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		})
	}

	errs.ResponseSuccessWithData(c, items)
}

// GetRolePermissions 获取角色的权限
func (h *RBACHandler) GetRolePermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid role ID")
		return
	}

	perms, err := h.rbacSvc.GetPermissionsByRoleID(c.Request.Context(), int64(id))
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	type PermItem struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}

	items := make([]*PermItem, 0, len(perms))
	for _, p := range perms {
		items = append(items, &PermItem{ID: p.ID, Name: p.Name, Code: p.Code})
	}

	errs.ResponseSuccessWithData(c, items)
}

// SetRolePermissions 设置角色权限
func (h *RBACHandler) SetRolePermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid role ID")
		return
	}

	var req struct {
		PermissionIDs []int64 `json:"permission_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, err.Error())
		return
	}

	if err := h.rbacSvc.SetRolePermissions(c.Request.Context(), int64(id), req.PermissionIDs); err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	errs.ResponseSuccess(c)
}
