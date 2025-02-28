package main

import (
	"context"
	"log"

	"github.com/yourusername/project/internal/api"
	"github.com/yourusername/project/internal/config"
	"github.com/yourusername/project/internal/middleware"
	"go.uber.org/fx"
)

// @title           Project API
// @version         1.0
// @description     Project API Documentation
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	app := fx.New(
		// 의존성 제공
		fx.Provide(
			config.NewConfig,
			middleware.NewLoggerMiddleware,
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
