package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/api"
	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/internal/middleware"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/pkg/database"
	"github.com/gudcks0305/payments-apply/pkg/logger"
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

				// 라우트 설정
				p.Router.SetupRoutes(p.Engine)
				// 정적 파일 제공 (프로덕션 환경에서 클라이언트 앱 서빙)
				/*p.Engine.Static("/assets", "./client/dist/assets")
				p.Engine.StaticFile("/", "./client/dist/index.html")

				// 모든 경로를 SPA로 리다이렉트 (클라이언트 라우팅 지원)
				p.Engine.NoRoute(func(c *gin.Context) {
					c.File("./client/dist/index.html")
				})*/
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
			config.NewConfig,
			logger.NewLogger,
			portone.NewClient,
			database.NewDatabase,
			api.NewHandler,
			newGinEngine,
		),
		// 애플리케이션 실행
		fx.Invoke(startServer),
	)
	app.Run()
}

func newGinEngine() *gin.Engine {
	r := gin.Default()
	// CORS 미들웨어 설정 추가
	r.Use(middleware.SetupCORS())
	return r
}
