package storage

import "context"

type IDB interface {
	RegisterMetrics(dbname string) (err error)
	Close() (err error)
	Migrate() (err error)

	GetFirst(ctx context.Context, dest interface{}, conds ...interface{}) (err error)
	GetByScript(ctx context.Context, dest interface{}, sql string, value []interface{}, limit, offset int) (err error)
}
