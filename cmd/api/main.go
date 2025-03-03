package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/gudcks0305/payments-apply/docs" // Swagger 문서를 임포트합니다
	"github.com/gudcks0305/payments-apply/internal/api"
	"github.com/gudcks0305/payments-apply/internal/api/handler"
	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/internal/middleware"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/internal/repository"
	"github.com/gudcks0305/payments-apply/internal/service"
	"github.com/gudcks0305/payments-apply/pkg/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type HandlerParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *config.Config
	DB        *gorm.DB
	Engine    *gin.Engine
	Router    *api.Handler
}

func startServer(p HandlerParams) {
	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// 설정에 따라 데이터베이스 마이그레이션 실행
				if p.Config.Database.AutoMigrate {
					if err := database.DBAutoMigrate(p.DB); err != nil {
						return err
					}
				}
				p.Engine.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

				go p.Engine.Run(":" + p.Config.Server.Port)

				return nil
			},
		},
	)
}

// @title           Payments API
// @version         1.0
// @description     Payments API Documentation
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	app := fx.New(
		// 의존성 제공
		fx.Provide(
			// etc ...
			config.NewConfig,
			portone.NewClient,
			database.NewDatabase,
			// svc
			service.NewPaymentService,
			//repo
			repository.NewPaymentRepository,

			// handler...
			handler.NewPaymentHandler,
			api.NewHandler,

			newGinEngine,
		),
		// 애플리케이션 실행
		fx.Invoke(startServer),
	)
	app.Run()
}

func newGinEngine(config *config.Config) *gin.Engine {
	r := gin.Default()
	// CORS 미들웨어 설정 추가
	gin.SetMode(config.Server.Mode)
	r.Use(middleware.SetupCORS())
	return r
}
