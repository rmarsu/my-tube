package server

import (
	"context"
	"fmt"
	"myTube/pkg/log"
	"net/http"
	"time"
)

type Server struct {
	http.Server
}

func NewServer(addr string, handler *http.ServeMux) *Server {
	return &Server{http.Server{Addr: fmt.Sprintf(":%s", addr), Handler: handler}}
}

func (s *Server) Start(ctx context.Context) error {
	closer := &Closer{}

	closer.Add(s.Shutdown)

	closer.Add(func(ctx context.Context) error {
		time.Sleep(5 * time.Second)
		return nil
	})

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("Error starting server: %v", err))
		}
	}()

	log.Info(fmt.Sprintf("listening on %s...\n", s.Addr))

	<-ctx.Done()

	log.Info("shutting down server...\n")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := closer.Close(shutdownCtx); err != nil {
		log.Error(fmt.Sprintf("Error shutting down server: %v\n", err))
	}
	return nil
}
