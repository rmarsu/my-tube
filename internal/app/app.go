package app

import (
	"context"
	"myTube/internal/config"
	"myTube/internal/repository"
	"myTube/internal/server"
	"myTube/internal/service"
	transport "myTube/internal/transport/rest"
	database "myTube/pkg/database/postgres"
	"myTube/pkg/hash"
	"myTube/pkg/log"
	"os/signal"
	"syscall"
)

func Run() {
	config := config.LoadConfig("configs/config.yaml")

	db , err := database.Connect()
	if err!= nil {
          log.Fatal(err)
     }
	defer db.Close(context.Background())

	repos := repository.NewRepositories(db)

	//deps
	hasher := hash.NewSHA1Hasher(config.Salt)

	services := service.NewServices(&service.Deps{
		Repos: repos,
		Hasher: hasher,

	})
	handler := transport.NewHandler(services)
	server := server.NewServer(config.Port, handler.InitRoutes())
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	if err := server.Start(ctx); err!= nil {
		log.Fatal(err)
     }
}
