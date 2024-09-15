package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sabyradinov/go-dating-matchmaking/config"
	"net/http"
	"strconv"
	"time"
)

type Server interface {
	Start() (serverChannel chan error)
	Stop() error
}

// server structure for web server
type server struct {
	srv           *http.Server
	srvCh         chan error
	stopTimeoutMS time.Duration
}

// Start start web server
func (s *server) Start() (serverChannel chan error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.srvCh <- err
		}
	}()

	return s.srvCh
}

// Stop stop web server
func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeoutMS)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}
	return nil
}

// New constructor for web server
func New(cfg *config.Configs, gin *Router) (Server, error) {
	if cfg.WebServer.Port <= 0 {
		return nil, fmt.Errorf("bad port value %d for server", cfg.WebServer.Port)
	}

	if cfg.StopTimeoutMS < 1 {
		return nil, fmt.Errorf("bad stop timeout %d for server", cfg.StopTimeoutMS)
	}

	s := &server{
		srv: &http.Server{
			Addr:    ":" + strconv.Itoa(cfg.WebServer.Port),
			Handler: gin.Engine,
		},
		stopTimeoutMS: time.Duration(cfg.StopTimeoutMS) * time.Millisecond,
	}

	return s, nil
}
