package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/api/handler"
	"gorm.io/gorm"
)

// @title 결제 시스템 API
// @version 1.0
// @description 결제 시스템을 위한 REST API 서비스
// @BasePath /api/v1

type Handler struct {
	db             *gorm.DB
	paymentHandler *handler.PaymentHandler
}

func NewHandler(db *gorm.DB, engine *gin.Engine, paymentHandler *handler.PaymentHandler) *Handler {
	h := &Handler{db: db, paymentHandler: paymentHandler}
	h.SetupRoutes(engine)
	return h
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		payments := api.Group("/payments")
		{
			payments.GET("", func(context *gin.Context) {
				context.JSON(200, gin.H{"message": "POST /api/v1/payments"})
			})
			payments.POST("", h.paymentHandler.CreatePayment)
			payments.PUT("/:id/complete", h.paymentHandler.ConfirmWithCompletePayment)
			payments.PUT("/complete", h.paymentHandler.ConfirmWithCompletePaymentBasic)
			payments.GET("/imp/:impUID", h.paymentHandler.GetPaymentByImpUID)
			payments.POST("/imp/:impUID/cancel", h.paymentHandler.CancelPaymentByImpUID)
		}

	}

}
