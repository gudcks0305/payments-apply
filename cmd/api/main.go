package main

import (
	"context"
	"log"

	"github.com/gudcks0305/payments-apply/internal/api"
	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/internal/middleware"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/pkg/logger"
	"go.uber.org/fx"
)

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
			middleware.NewLoggerMiddleware,
			logger.NewLogger,
			portone.NewClient,
			api.NewRouter,
		),
		// 애플리케이션 실행
		fx.Invoke(func(router *api.Router) {
			router.Run()
		}),
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
