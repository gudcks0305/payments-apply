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
	ID          uuid.UUID `gorm:"primaryKey;type:varchar(36)"` // MerchantUID
	ImpUID      string    `gorm:"index"`
	ProductName string
	Amount      int
	Status      PaymentStatusType `gorm:"default:pending"`
	PayMethod   string            `gorm:"default:card"`
	ErrCode     string            `gorm:"default:"`
	ErrMessage  string            `gorm:"default:"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
