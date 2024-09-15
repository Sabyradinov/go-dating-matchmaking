package gorm

import (
	"context"
	"time"
)

func (dbm *dbClient) GetFirst(ctx context.Context, dest interface{}, conds ...interface{}) (err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := dbm.db.WithContext(ctxTimeout)
	err = tx.First(&dest, conds).Error
	return
}

func (dbm *dbClient) GetByScript(ctx context.Context, dest interface{}, sql string, value []interface{}, limit, offset int) (err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := dbm.db.WithContext(ctxTimeout)
	err = tx.Raw(sql, value...).Limit(limit).Offset(offset).Scan(&dest).Error

	return
}
