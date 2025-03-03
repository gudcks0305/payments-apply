package handler

import (
	"fmt"
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
		c.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	payment, err := ph.paymentService.CreatePayment(&paymentRequest)
	if err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	c.JSON(201, dto.APIResponseCreated(payment))
}

func (ph *PaymentHandler) ConfirmWithCompletePayment(c *gin.Context) {
	id := c.Param("id")
	var paymentData *portone.PaymentClientResponse
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	payment, err := ph.paymentService.ConfirmWithCompletePayment(id, paymentData)
	if err != nil {
		fmt.Println(err)
		appErr := errors.MapError(err)
		c.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	c.JSON(200, dto.APIResponseSuccess(payment))
}

func (ph *PaymentHandler) GetPaymentByImpUID(context *gin.Context) {
	impUID := context.Param("impUID")
	payment, err := ph.paymentService.GetPaymentByIMPUID(impUID)
	if err != nil {
		appErr := errors.MapError(err)
		context.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	context.JSON(200, dto.APIResponseSuccess(payment))
}

func (ph *PaymentHandler) CancelPaymentByImpUID(context *gin.Context) {
	impUID := context.Param("impUID")
	payment, err := ph.paymentService.CancelPaymentByIMPUID(impUID)
	if err != nil {
		appErr := errors.MapError(err)
		context.JSON(appErr.StatusCode, dto.APIResponseError[string](appErr))
		return
	}

	context.JSON(200, dto.APIResponseSuccess(payment))
}
