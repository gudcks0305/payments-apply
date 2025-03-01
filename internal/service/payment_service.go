package service

import (
	"github.com/gudcks0305/payments-apply/internal/repository"
)

type PaymentService struct {
	repository *repository.PaymentRepository
}

func NewPaymentService(repository *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repository: repository}
}
