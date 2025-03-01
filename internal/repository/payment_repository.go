package repository

import (
	"github.com/google/uuid"
	"github.com/gudcks0305/payments-apply/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (pr *PaymentRepository) CreatePayment(payment *model.Payment) error {
	return pr.db.Create(payment).Error
}

func (pr *PaymentRepository) GetPaymentByID(id uuid.UUID) (*model.Payment, error) {
	var payment model.Payment
	err := pr.db.First(&payment, id).Error
	return &payment, err
}
