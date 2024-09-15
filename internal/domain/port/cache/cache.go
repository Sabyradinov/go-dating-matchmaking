package cache

import (
	"context"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) (res string, err error)
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) (err error)
	Exists(ctx context.Context, key string) (res bool)
	Del(ctx context.Context, key string) (err error)
	RegisterMetrics()
}
