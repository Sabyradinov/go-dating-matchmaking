package gorm

import (
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/entity"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// dbClient instance to db client
type dbClient struct {
	db  *gorm.DB
	cfg *config.Configs
}

// InitDB initialize db
func InitDB(cfg *config.Configs) (storage.IDB, error) {

	gormDB, err := gorm.Open(postgres.Open(cfg.DB.ConnectionString), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	return &dbClient{db: gormDB}, err
}

// Close db client close
func (dbm *dbClient) Close() (err error) {
	db, err := dbm.db.DB()
	if err != nil {
		return
	}

	err = db.Close()
	return
}

// RegisterMetrics register prometheus metrics
func (dbm *dbClient) RegisterMetrics(dbname string) (err error) {

	return
}

func (dbm *dbClient) Migrate() (err error) {

	// Create the database if it doesn't exist
	err = dbm.db.Exec("CREATE DATABASE IF NOT EXISTS datingapp").Error
	if err != nil {
		return err
	}

	// Create table for `User`
	err = dbm.db.Migrator().CreateTable(&entity.User{})
	if err != nil {
		return
	}

	// Create table for `Preference`
	err = dbm.db.Migrator().CreateTable(&entity.Preference{})
	if err != nil {
		return
	}
	// Create index for Gender field
	err = dbm.db.Migrator().CreateIndex(&entity.User{}, "gender")
	if err != nil {
		return err
	}

	// Create index for Age field
	err = dbm.db.Migrator().CreateIndex(&entity.User{}, "age")
	if err != nil {
		return err
	}

	// Create index for Location field
	err = dbm.db.Migrator().CreateIndex(&entity.User{}, "location")
	if err != nil {
		return err
	}

	// Create index for LastActive field
	err = dbm.db.Migrator().CreateIndex(&entity.User{}, "last_active")
	if err != nil {
		return err
	}

	return
}
