package gorm

import (
	"context"
	"time"
)

func (dbm *dbClient) GetFirst(ctx context.Context, dest interface{}, query interface{}, args ...interface{}) (err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := dbm.db.WithContext(ctxTimeout)
	err = tx.Where(query, args).First(&dest).Error
	return
}

func (dbm *dbClient) GetByScript(ctx context.Context, dest interface{}, limit, offset int, sql string, values []interface{}) (err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := dbm.db.WithContext(ctxTimeout)
	err = tx.Raw(sql, values...).Limit(limit).Offset(offset).Find(dest).Error

	return
}
