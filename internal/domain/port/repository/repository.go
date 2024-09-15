package repository

import (
	"context"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/entity"
)

type IUser interface {
	GetById(ctx context.Context, userID string) (res entity.User, err error)
	GetPotentialMatches(ctx context.Context, currentUser entity.User, offset, limit int) ([]entity.User, error)
}
