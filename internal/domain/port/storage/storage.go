package storage

import "context"

type IDB interface {
	RegisterMetrics(dbname string) (err error)
	Close() (err error)
	Migrate() (err error)

	GetFirst(ctx context.Context, dest interface{}, query interface{}, args ...interface{}) (err error)
	GetByScript(ctx context.Context, dest interface{}, limit, offset int, sql string, values []interface{}) (err error)
}
