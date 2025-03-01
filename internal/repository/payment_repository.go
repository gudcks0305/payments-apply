package repository

import (
	"github.com/google/uuid"
	"github.com/gudcks0305/payments-apply/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (pr *PaymentRepository) CreatePayment(payment *model.Payment) error {
	return pr.DB.Create(payment).Error
}

func (pr *PaymentRepository) GetPaymentByID(id uuid.UUID) (*model.Payment, error) {
	var payment model.Payment
	err := pr.DB.First(&payment, id).Error
	return &payment, err
}
