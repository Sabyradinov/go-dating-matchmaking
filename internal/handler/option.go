package handler

import (
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/logger"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/service"
)

type Options struct {
	Services *service.Builder
	Logger   logger.AppLogger
	Cfg      *config.Configs
}
