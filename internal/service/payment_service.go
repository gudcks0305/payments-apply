package service

import (
	"github.com/google/uuid"
	"github.com/gudcks0305/payments-apply/internal/dto"
	"github.com/gudcks0305/payments-apply/internal/model"
	"github.com/gudcks0305/payments-apply/internal/repository"
)

type PaymentService struct {
	repository *repository.PaymentRepository
}

func (s PaymentService) CreatePayment(d *dto.PaymentCreateRequest) (*dto.IdResponse[uuid.UUID], error) {
	payment := &model.Payment{
		Amount:      d.Amount,
		PayMethod:   d.PayMethod,
		ProductName: d.ProductName,
		Status:      model.StatusPending,
	}
	err := s.repository.CreatePayment(payment)
	if err != nil {
		return nil, err
	}
	return &dto.IdResponse[uuid.UUID]{ID: payment.ID}, nil
}

func NewPaymentService(repository *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repository: repository}
}
