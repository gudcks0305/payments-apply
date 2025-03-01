package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/api/handler"
	"gorm.io/gorm"
)

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
	// API 라우트 설정
	api := r.Group("/api/v1")
	{
		// Payment 라우트 설정
		payments := api.Group("/payments")
		{
			payments.GET("", func(context *gin.Context) {
				context.JSON(200, gin.H{"message": "POST /api/v1/payments"})
			})
			payments.POST("", h.paymentHandler.CreatePayment)
			payments.PUT("/:id/complete", h.paymentHandler.ConfirmWithCompletePayment)
			payments.GET("/imp/:impUID", h.paymentHandler.GetPaymentByImpUID)
		}

	}

}
