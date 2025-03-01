package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/dto"
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
		//TODO error 나중에 한번에 Mapping 화 시키기
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment, err := ph.paymentService.CreatePayment(&paymentRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, dto.APIResponseCreated(payment))
}

func (ph *PaymentHandler) ConfirmWithCompletePayment(c *gin.Context) {
	var paymentData dto.PaymentData
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment, err := ph.paymentService.ConfirmWithCompletePayment(&paymentData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dto.APIResponseSuccess(payment))
}

func (ph *PaymentHandler) GetPaymentByImpUID(context *gin.Context) {
	impUID := context.Param("impUID")
	payment, err := ph.paymentService.GetPaymentByIMPUID(impUID)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, dto.APIResponseSuccess(payment))
}
