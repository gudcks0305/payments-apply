package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrPaymentNotFound     = errors.New("payment not found")
	ErrPaymentFailed       = errors.New("payment processing failed")
	ErrPaymentAlreadyExist = errors.New("payment with this merchant_uid already exists")
	ErrInvalidPaymentState = errors.New("invalid payment state for this operation")
	ErrInvalidAmount       = errors.New("invalid payment amount")
	ErrInvalidPaymentInfo  = errors.New("invalid payment information")
	ErrPaymentCanceled     = errors.New("payment was canceled")
	ErrPortOneError        = errors.New("portone service error")
	ErrPaymentExpired      = errors.New("payment request has expired")
)

type AppError struct {
	Err        error
	StatusCode int
	Message    string
}

func (e AppError) Error() string {
	return e.Message
}

func MapError(err error) AppError {
	switch {
	// 결제 관련 에러 처리 추가
	case errors.Is(err, ErrPaymentNotFound):
		return AppError{
			Err:        err,
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		}
	case errors.Is(err, ErrPaymentAlreadyExist):
		return AppError{
			Err:        err,
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
		}
	case errors.Is(err, ErrInvalidAmount), errors.Is(err, ErrInvalidPaymentInfo),
		errors.Is(err, ErrInvalidPaymentState):
		return AppError{
			Err:        err,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	case errors.Is(err, ErrPaymentFailed), errors.Is(err, ErrPortOneError):
		return AppError{
			Err:        err,
			StatusCode: http.StatusServiceUnavailable,
			Message:    err.Error(),
		}
	case errors.Is(err, ErrPaymentCanceled):
		return AppError{
			Err:        err,
			StatusCode: http.StatusOK,
			Message:    err.Error(),
		}
	case errors.Is(err, ErrPaymentExpired):
		return AppError{
			Err:        err,
			StatusCode: http.StatusGone,
			Message:    err.Error(),
		}
	default:
		// 바인딩 에러 처리
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var messages []string
			for _, e := range validationErrors {
				messages = append(messages, fmt.Sprintf("Field '%s' %s", e.Field(), getValidationErrorMsg(e)))
			}
			return AppError{
				Err:        err,
				StatusCode: http.StatusBadRequest,
				Message:    strings.Join(messages, "; "),
			}
		}

		// 기타 바인딩 에러
		if strings.Contains(err.Error(), "binding") {
			return AppError{
				Err:        err,
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid request format",
			}
		}

		fmt.Println("[ERROR]", err)
		return AppError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
}

// 유효성 검사 에러 메시지 생성
func getValidationErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "min":
		return fmt.Sprintf("must be at least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", e.Param())
	case "email":
		return "must be a valid email address"
	case "gt":
		return fmt.Sprintf("must be greater than %s", e.Param())
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", e.Param())
	default:
		return fmt.Sprintf("failed on '%s' validation", e.Tag())
	}
}
