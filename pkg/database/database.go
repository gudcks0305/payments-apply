package database

import (
	"github.com/gudcks0305/payments-apply/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const SupportedDrivers = "sqlite3"

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	//driver := cfg.Database.Driver
	//if driver != SupportedDrivers {
	//	panic("지원하지 않는 데이터베이스 드라이버입니다")
	//}
	if cfg.Database.Host == ":memory:" {
		return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}

	return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

}
