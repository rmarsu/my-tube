package app

import (
	"context"
	"myTube/internal/config"
	"myTube/internal/migrations"
	"myTube/internal/repository"
	"myTube/internal/server"
	"myTube/internal/service"
	transport "myTube/internal/transport/rest"
	"myTube/pkg/auth"
	database "myTube/pkg/database/postgres"
	"myTube/pkg/hash"
	"myTube/pkg/log"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	config := config.LoadConfig("configs/config.yaml")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = migrations.InitTables(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repos := repository.NewRepositories(db)

	//deps
	hasher := hash.NewSHA256Hasher(config.Salt)
	tokenManager, err := auth.NewManager(config.JWTSecret)
	if err != nil {
		log.Fatal(err)
	}

	services := service.NewServices(&service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AccessTokenTTL:  time.Hour * 2,
		RefreshTokenTTL: time.Hour * 720,
		// Add other dependencies as needed

	})
	handler := transport.NewHandler(services)
	server := server.NewServer(config.Port, handler.InitRoutes())
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := server.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
