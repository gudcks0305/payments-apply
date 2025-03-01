package model

import "time"

type PaymentStatusType string

const (
	StatusPending   PaymentStatusType = "pending"
	StatusCompleted PaymentStatusType = "completed"
	StatusFailed    PaymentStatusType = "failed"
	StatusCancelled PaymentStatusType = "cancelled"
)

type Payment struct {
	ID           uint   `gorm:"primaryKey"`
	MerchantUID  string `gorm:"uniqueIndex"`
	CustomerID   string `gorm:"index"`
	Amount       int
	Status       PaymentStatusType
	PayMethod    string
	PaidAt       *time.Time
	FailReason   string
	CancelReason string
	CancelledAt  *time.Time
	RequestData  string `gorm:"type:text"`
	ResponseData string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
