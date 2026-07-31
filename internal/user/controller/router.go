package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/user/handler"
	medHandler "github.com/teakingwang/backend/internal/user/medical/handler"
	medSvc "github.com/teakingwang/backend/internal/user/medical/service"
	"github.com/teakingwang/backend/internal/user/s
	"github.com/teakingwang/backend/pkg/middleware"
)

func NewHTTPRouter(userSvc service.UserService, rbacSvc service.RBACService, medicalSvc medSvc.MedicalRecordService) *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 32 << 20 // 32MB

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 静态文件 - 不需要JWT
	r.Static("/uploads", "./uploads")

	// JWT
	r.Use(middleware.JWTGinMiddleware())

	userHandler := handler.NewUserHandler(userSvc, rbacSvc)
	rbacHandler := handler.NewRBACHandler(rbacSvc)
	medH := medHandler.NewMedicalHandler(medicalSvc)

	// 用户路由
	v1 := r.Group("/v1/user")
	{
		v1.POST("/login", userHandler.Login)
		v1.GET("/list", userHandler.GetUserList)
		v1.GET("/:id", userHandler.GetUser)
	}

	v1Admin := r.Group("/v1/user")
	{
		v1Admin.POST("/create", middleware.RequirePermission("user:create"), userHandler.CreateUser)
		v1Admin.PUT("/:id", middleware.RequirePermission("user:update"), userHandler.UpdateUser)
		v1Admin.DELETE("/:id", middleware.RequirePermission("user:delete"), userHandler.DeleteUser)
	}

	// RBAC 路由
	rbacGroup := r.Group("/v1/rbac")
	{
		rbacGroup.GET("/roles", rbacHandler.GetRoles)
		rbacGroup.GET("/permissions", rbacHandler.GetPermissions)
		rbacGroup.GET("/roles/:id/permissions", rbacHandler.GetRolePermissions)
	}

	rbacAdminGroup := r.Group("/v1/rbac")
	rbacAdminGroup.Use(middleware.RequirePermission("role:manage"))
	{
		rbacAdminGroup.PUT("/roles/:id/permissions", rbacHandler.SetRolePermissions)
	}

	// 医疗档案路由
	medGroup := r.Group("/v1/medical")
	{
		medGroup.GET("/list", medH.ListRecords)
		medGroup.GET("/:id", medH.GetRecord)
		medGroup.POST("/create", medH.CreateRecord)
		medGroup.PUT("/:id", medH.UpdateRecord)
		medGroup.DELETE("/:id", medH.DeleteRecord)
		medGroup.DELETE("/invoice/:id", medH.DeleteInvoice)
	}

	return r
}
