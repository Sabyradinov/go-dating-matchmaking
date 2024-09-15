package main

import (
	"log"

	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/gorm"
)

func main() {
	cfg, err := config.Init("config.json")
	if err != nil {
		log.Fatalf("config init error: %v", err)
	}

	db, err := gorm.InitDB(cfg)
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf("db migrate error: %v", err)
	}

	defer db.Close()

}
