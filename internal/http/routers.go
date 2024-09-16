package http

import (
	"github.com/Sabyradinov/go-dating-matchmaking/cmd/docs"
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/logger"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func NewRouter(cfg *config.Configs, log logger.AppLogger, services *service.Builder) (*Router, error) {
	//create gin instance
	router := gin.New()
	//create handlers
	handler := newHandler(cfg, log, services)

	//config gin logger
	if cfg.WebServer.GIN.UseLogger {
		router.Use(gin.Logger())
	}
	//config panic handler
	if cfg.WebServer.GIN.UseRecovery {
		router.Use(gin.CustomRecovery(log.HttpPanicHandler))
	}

	//config routes
	rg := router.Group("/api")
	{
		match := rg.Group("/match")
		{
			match.GET("/recommendations", handler.matching.GetPotentialMatches)
		}
	}

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// config swagger ui
	docs.SwaggerInfo.Host = cfg.SwaggerUI.Host
	docs.SwaggerInfo.Description = cfg.SwaggerUI.Description
	docs.SwaggerInfo.Title = cfg.SwaggerUI.PageTitle
	docs.SwaggerInfo.Schemes = cfg.SwaggerUI.Schemes

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{router}, nil
}
