package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/dto"
	"github.com/gudcks0305/payments-apply/internal/errors"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/internal/service"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (ph *PaymentHandler) CreatePayment(c *gin.Context) {
	var paymentRequest dto.PaymentCreateRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	payment, err := ph.paymentService.CreatePayment(&paymentRequest)
	if err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(201, dto.APIResponseCreated(payment))
}

func (ph *PaymentHandler) ConfirmWithCompletePayment(c *gin.Context) {
	var paymentData *portone.PaymentData
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	payment, err := ph.paymentService.ConfirmWithCompletePayment(paymentData)
	if err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(200, dto.APIResponseSuccess(payment))
}

func (ph *PaymentHandler) GetPaymentByImpUID(context *gin.Context) {
	impUID := context.Param("impUID")
	payment, err := ph.paymentService.GetPaymentByIMPUID(impUID)
	if err != nil {
		appErr := errors.MapError(err)
		context.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	context.JSON(200, dto.APIResponseSuccess(payment))
}
