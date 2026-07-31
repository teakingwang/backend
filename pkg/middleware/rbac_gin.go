package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/user/repository"
	"github.com/teakingwang/backend/pkg/db"
)

// RequirePermission 检查当前用户是否拥有指定权限
func RequirePermission(permCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No user context"})
			return
		}

		userIDStr, ok := userIDVal.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid user context"})
			return
		}

		// 简单检查：系统管理员（is_system）跳过权限检查
		// 也可以从 RBAC 仓库检查
		rbacRepo := repository.NewRBACRepo(db.GormDB)

		// 将 userID 转换为 int64
		var userID int64
		for _, ch := range userIDStr {
			userID = userID*10 + int64(ch-'0')
		}

		hasPermission, err := rbacRepo.CheckUserPermission(c.Request.Context(), userID, permCode)
		if err != nil || !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}

		c.Next()
	}
}
