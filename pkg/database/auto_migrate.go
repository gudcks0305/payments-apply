package database

import (
	"github.com/gudcks0305/payments-apply/internal/model"
	"gorm.io/gorm"
)

func DBAutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Payment{},
	)

}
