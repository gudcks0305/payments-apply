package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/project/internal/config"
	"github.com/yourusername/project/internal/middleware"
)

type Router struct {
	config *config.Config
	engine *gin.Engine
	logger *middleware.LoggerMiddleware
}

func NewRouter(
	config *config.Config,
	logger *middleware.LoggerMiddleware,
) *Router {
	engine := gin.New()

	// 미들웨어 설정
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger())

	// 서버 모드 설정
	gin.SetMode(config.Server.Mode)

	return &Router{
		config: config,
		engine: engine,
		logger: logger,
	}
}

func (r *Router) Run() {

	// Swagger 문서 설정
	//r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 서버 시작
	r.engine.Run(":" + r.config.Server.Port)
}
