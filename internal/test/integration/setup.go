package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/api"
	"github.com/gudcks0305/payments-apply/internal/api/handler"
	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/internal/middleware"
	"github.com/gudcks0305/payments-apply/internal/repository"
	"github.com/gudcks0305/payments-apply/internal/service"
	"github.com/gudcks0305/payments-apply/internal/test/mock"
	"github.com/gudcks0305/payments-apply/pkg/database"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func SetupGinApp(t *testing.T) (*gin.Engine, *api.Handler, *fxtest.App) {
	// 테스트 모드 설정
	gin.SetMode(gin.TestMode)

	// 테스트 의존성 준비
	var engine *gin.Engine
	var apiHandler *api.Handler

	// 테스트 앱 설정
	app := fxtest.New(
		t,
		fx.Provide(
			func() *config.Config {
				return &config.Config{
					Server: struct {
						Port string `mapstructure:"port"`
						Mode string `mapstructure:"mode"`
					}{
						Mode: "test",
						Port: "8080",
					},
					Database: struct {
						Host        string `mapstructure:"host"`
						Port        string `mapstructure:"port"`
						User        string `mapstructure:"user"`
						Password    string `mapstructure:"password"`
						DBName      string `mapstructure:"db_name"`
						Driver      string `mapstructure:"driver"`
						AutoMigrate bool   `mapstructure:"auto_migrate"`
					}{
						Driver:      "sqlite3",
						Host:        ":memory:", // Connection이 아닌 Host 사용
						AutoMigrate: true,
					},
					PortOne: struct {
						ImpKey    string `mapstructure:"imp_key"`
						ImpSecret string `mapstructure:"imp_secret"`
						BaseURL   string `mapstructure:"base_url"`
					}{
						ImpKey:    "test_imp_key",
						ImpSecret: "test_secret_key",
						BaseURL:   "https://api.iamport.kr",
					},
				}
			},
			database.NewDatabase,
			mock.NewMockClient,
			repository.NewPaymentRepository,
			service.NewPaymentService,
			handler.NewPaymentHandler,
			api.NewHandler,
			func() *gin.Engine {
				r := gin.New()
				r.Use(gin.Recovery())
				r.Use(middleware.SetupCORS())
				return r
			},
		),
		fx.Invoke(database.DBAutoMigrate),
		fx.Populate(&engine, &apiHandler),
	)

	return engine, apiHandler, app
}
