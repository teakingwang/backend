package handler

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserItem struct {
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Status     int8   `json:"status"`
	StatusName string `json:"status_name"`
	RoleID     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
}

type LoginResp struct {
	Token string    `json:"token"`
	User  *UserItem `json:"user"`
}

type GetUserListReq struct{}

type GetUserListResp struct {
	List  []*UserItem `json:"list"`
	Total int64       `json:"total"`
}

type GetUserResp struct {
	User *UserItem `json:"user"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
	Nickname string `json:"nickname" binding:"required"`
	RoleID   int64  `json:"role_id"`
}

type CreateUserResp struct {
	User *UserItem `json:"user"`
}

type UpdateUserReq struct {
	Password string `json:"password" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
	Nickname string `json:"nickname" binding:"omitempty"`
	RoleID   int64  `json:"role_id"`
	Status   *int8  `json:"status"`
}

type UpdateUserResp struct {
	User *UserItem `json:"user"`
}
