package service

import (
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/repo"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/service"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/service/matching"
)

type Builder struct {
	Matching service.IMatching
}

func Init(repo *repo.Builder) *Builder {
	return &Builder{
		Matching: matching.New(repo),
	}
}
