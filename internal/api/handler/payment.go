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

// CreatePayment godoc
// @Summary 결제 생성
// @Description 새로운 결제 정보를 생성합니다
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body dto.PaymentCreateRequest true "결제 생성 요청 데이터"
// @Router /payments [post]
// @Success 201 {object} dto.APIResponse[IdResponse[string]] "Created"
// @Failure 400 {object} dto.APIResponse[string] "Bad Request"
// @Failure 500 {object} dto.APIResponse[string] "Internal Server Error"
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

// ConfirmWithCompletePayment godoc
// @Summary 결제 완료 확인
// @Description 지정된 ID의 결제를 완료 상태로 변경합니다
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "결제 ID"
// @Param payment body portone.PaymentClientResponse true "결제 확인 데이터"
// @Router /payments/{id}/complete [put]
// @Success 200 {object} dto.APIResponse[portone.PaymentData] "Success"
// @Failure 400 {object} dto.APIResponse[string] "Bad Request"
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

// GetPaymentByImpUID godoc
// @Summary ptone UID로 결제 조회
// @Description ptone UID를 사용하여 결제 정보를 조회합니다
// @Tags payments
// @Accept json
// @Produce json
// @Param impUID path string true "ptone UID"
// @Router /payments/imp/{impUID} [get]
// @Success 200 {object} dto.APIResponse[portone.PaymentData] "Success"
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

// CancelPaymentByImpUID godoc
// @Summary ptone UID로 결제 취소
// @Description ptone UID를 사용하여 결제를 취소합니다
// @Tags payments
// @Accept json
// @Produce json
// @Param impUID path string true "ptone UID"
// @Router /payments/imp/{impUID}/cancel [post]
// @Success 200 {object} dto.APIResponse[portone.PaymentData] "Success"
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
