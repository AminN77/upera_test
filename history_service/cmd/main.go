package main

import (
	"github.com/AminN77/upera_test/history_service/api/controller"
	"github.com/AminN77/upera_test/history_service/cmd/setup"
	"github.com/AminN77/upera_test/history_service/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	repo := internal.NewMongoRepository()
	subscriber := internal.NewKafkaEventSubscriber(repo)
	if err := subscriber.Subscribe(); err != nil {
		log.Fatal(err)
	}

	// service & controller
	srv := internal.NewService(repo)
	con := controller.NewController(srv)

	router := setup.SetRouter(con)

	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		log.Fatal(err)
	}
}
