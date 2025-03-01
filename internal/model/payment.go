package model

import (
	"gorm.io/gorm"
	"time"
)
import "github.com/google/uuid"

type PaymentStatusType string

const (
	StatusPending   PaymentStatusType = "pending"
	StatusCompleted PaymentStatusType = "completed"
	StatusFailed    PaymentStatusType = "failed"
	StatusCancelled PaymentStatusType = "cancelled"
)

type Payment struct {
	ID           uuid.UUID `gorm:"primaryKey;type:varchar(36)"` // MerchantUID
	ImpUID       string    `gorm:"index"`
	ProductName  string
	Amount       int
	Status       PaymentStatusType
	PayMethod    string
	PaidAt       *time.Time
	ErrorCode    string
	ErrorMessage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
