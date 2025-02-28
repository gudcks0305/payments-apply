package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/internal/middleware"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/pkg/logger"
)

type Router struct {
	config    *config.Config
	engine    *gin.Engine
	logger    *middleware.LoggerMiddleware
	portone   *portone.Client
	sysLogger logger.Logger
}

func NewRouter(
	config *config.Config,
	logger *middleware.LoggerMiddleware,
	portone *portone.Client,
	sysLogger logger.Logger,
) *Router {
	engine := gin.New()

	// 미들웨어 설정
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger())

	// 서버 모드 설정
	gin.SetMode(config.Server.Mode)

	return &Router{
		config:    config,
		engine:    engine,
		logger:    logger,
		portone:   portone,
		sysLogger: sysLogger,
	}
}

func (r *Router) Run() {
	// API 라우트 설정
	//api := r.engine.Group("/api/v1")

	// 서버 시작
	r.engine.Run(":" + r.config.Server.Port)
}
