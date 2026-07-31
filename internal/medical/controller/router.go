package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/backend/internal/medical/handler"
	"github.com/teakingwang/backend/internal/medical/service"
	"github.com/teakingwang/backend/pkg/middleware"
)

func NewHTTPRouter(svc service.MedicalRecordService, uploadDir string) *gin.Engine {
	r := gin.Default()

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

	r.Use(middleware.JWTGinMiddleware())

	// 静态文件服务
	r.Static("/uploads", uploadDir)

	h := handler.NewMedicalHandler(svc)

	v1 := r.Group("/v1/medical")
	{
		v1.GET("/list", h.ListRecords)
		v1.GET("/:id", h.GetRecord)
		v1.POST("/create", h.CreateRecord)
		v1.DELETE("/:id", h.DeleteRecord)
	}

	return r
}
