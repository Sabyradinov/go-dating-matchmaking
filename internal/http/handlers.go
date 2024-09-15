package http

import (
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/logger"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/service"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/handler"
)

type Handlers struct {
	matching *handler.Matching
}

func newHandler(cfg *config.Configs, logger logger.AppLogger, services *service.Builder) Handlers {
	return Handlers{
		matching: handler.NewMatching(&handler.Options{Cfg: cfg, Logger: logger, Services: services}),
	}
}
