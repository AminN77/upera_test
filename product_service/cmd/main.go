package main

import (
	"github.com/AminN77/upera_test/product_service/api/controller"
	"github.com/AminN77/upera_test/product_service/cmd/setup"
	"github.com/AminN77/upera_test/product_service/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// agent
	repo := internal.NewPostgresRepository()
	publisher := internal.NewKafkaEventPublisher()

	// service & controller
	srv := internal.NewService(repo, publisher)
	con := controller.NewController(srv)

	// setup router
	router := setup.SetRouter(con)

	// run
	if err := router.Listen(os.Getenv("API_PORT")); err != nil {
		log.Fatal(err)
	}
}
