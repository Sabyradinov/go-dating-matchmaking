package service

import (
	"context"
)

type IMatching interface {
	GetPotentialMatches(ctx context.Context, userId string, offset, limit int) (res interface{}, err error)
	GetUserById(ctx context.Context, userId string) (res interface{}, err error)
}
