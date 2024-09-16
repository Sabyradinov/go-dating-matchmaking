package main

import (
	"fmt"
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/logger"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/gorm"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/repo"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/service"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/http"
	"os"
	"os/signal"
	"syscall"
)

// @version 1.0
// @schemes https
// @BasePath /api
// @query.collection.format multi
func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Init("config.json")
	if err != nil {
		panic(fmt.Sprintf("config init error: %v", err))
	}

	db, err := gorm.InitDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("db init error: %v", err))
	}
	repos := repo.Init(db)
	srv := service.Init(repos)
	appLogger := logger.New()

	gin, err := http.NewRouter(cfg, appLogger, srv)
	if err != nil {
		panic(fmt.Sprintf("gin router init error: %v", err))
	}

	server, err := http.New(cfg, gin)
	if err != nil {
		panic(fmt.Sprintf("create server error: %v", err))
	}

	startServerError := server.Start()

	var stopReason string
	select {
	case err = <-startServerError:
		stopReason = fmt.Sprintf("start server error: %v", err)
	case qs := <-quit:
		stopReason = fmt.Sprintf("received signal %s", qs.String())
	}

	fmt.Printf("%s\nshutting down server...\n", stopReason)
	err = server.Stop()
	if err != nil {
		fmt.Printf("stop server error: %v\n", err)
		return
	}

	err = db.Close()
	if err != nil {
		fmt.Printf("stop db error: %v\n", err)
		return
	}

	fmt.Println("server stopped")
}
