package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/user/model"
	"github.com/teakingwang/backend/internal/user/service"
	"github.com/teakingwang/backend/pkg/errs"
	"github.com/teakingwang/backend/pkg/logger"
)

type UserHandler struct {
	svc     service.UserService
	rbacSvc service.RBACService
}

func NewUserHandler(svc service.UserService, rbacSvc service.RBACService) *UserHandler {
	return &UserHandler{svc: svc, rbacSvc: rbacSvc}
}

func dtoToItem(dto *service.UserDTO) *UserItem {
	if dto == nil {
		return nil
	}
	return &UserItem{
		UserID:     fmt.Sprintf("%d", dto.UserID),
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Email:      dto.Email,
		Mobile:     dto.Mobile,
		Status:     int8(dto.Status),
		StatusName: dto.StatusName,
		RoleID:     fmt.Sprintf("%d", dto.RoleID),
		RoleName:   dto.RoleName,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	req := &LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, err.Error())
		return
	}

	token, userDTO, err := h.svc.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		errs.ResponseError(c, errs.CodeUnauthorized, err.Error())
		return
	}

	resp := &LoginResp{
		Token: token,
		User:  dtoToItem(userDTO),
	}

	errs.ResponseSuccessWithData(c, resp)
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	req := GetUserListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, err.Error())
		return
	}

	userDTOList, total, err := h.svc.GetUserList(c.Request.Context())
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	list := make([]*UserItem, 0, len(userDTOList))
	for _, dto := range userDTOList {
		list = append(list, dtoToItem(dto))
	}

	resp := GetUserListResp{
		List:  list,
		Total: total,
	}

	errs.ResponseSuccessWithData(c, resp)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	logger.Info("GetUser called with ID:", id)
	if id == "" {
		errs.ResponseError(c, errs.CodeInvalidArgs, "missing user ID")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid user ID")
		return
	}

	userDTO, err := h.svc.GetUser(c.Request.Context(), idInt)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	if userDTO == nil {
		errs.ResponseError(c, errs.CodeNotFound, "用户不存在")
		return
	}

	resp := GetUserResp{
		User: dtoToItem(userDTO),
	}

	errs.ResponseSuccessWithData(c, resp)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := &CreateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, err.Error())
		return
	}

	userDTO, err := h.svc.CreateUser(c.Request.Context(), req.Username, req.Password, req.Email, req.Nickname, req.RoleID)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	resp := CreateUserResp{
		User: dtoToItem(userDTO),
	}

	errs.ResponseSuccessWithData(c, resp)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	logger.Info("DeleteUser called with ID:", id)
	if id == "" {
		errs.ResponseError(c, errs.CodeInvalidArgs, "missing user ID")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid user ID")
		return
	}

	err = h.svc.DeleteUser(c.Request.Context(), idInt)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	errs.ResponseSuccess(c)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	logger.Info("UpdateUser called with ID:", id)
	if id == "" {
		errs.ResponseError(c, errs.CodeInvalidArgs, "missing user ID")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, "invalid user ID")
		return
	}

	req := &UpdateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.ResponseError(c, errs.CodeInvalidArgs, err.Error())
		return
	}

	var status *model.UserStatus
	if req.Status != nil {
		s := model.UserStatus(*req.Status)
		status = &s
	}

	userDTO, err := h.svc.UpdateUser(c.Request.Context(), idInt, req.Password, req.Email, req.Nickname, req.RoleID, status)
	if err != nil {
		errs.ResponseError(c, errs.CodeDatabaseError, err.Error())
		return
	}

	resp := UpdateUserResp{
		User: dtoToItem(userDTO),
	}

	errs.ResponseSuccessWithData(c, resp)
}
