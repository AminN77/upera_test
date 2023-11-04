package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//repo := internal.NewMongoRepository()

	// service & controller
	//srv := internal.NewService(repo)
	//con := controller.NewController(srv)
	//
	//// setup router
	//router := setup.SetRouter(con)
	//
	//// run
	//if err := router.Listen(os.Getenv("API_PORT")); err != nil {
	//	panic(err)
	//}
}
