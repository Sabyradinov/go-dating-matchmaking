package service

import (
	"context"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/model"
)

type IMatching interface {
	GetPotentialMatches(ctx context.Context, userId string, offset, limit int) (res model.UserResponse, err error)
	GetUserById(ctx context.Context, userId string) (res interface{}, err error)
}
