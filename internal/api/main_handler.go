package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	// API 라우트 설정
	api := r.Group("/api/v1")
	{
		// Payment 라우트 설정
		api.POST("/payment", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "payment",
			})
		})

	}
}
