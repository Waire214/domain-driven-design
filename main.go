package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"presentation/interfaces"
	"presentation/routes"
	"presentation/sharedinfrastructure/persistence"
	"github.com/joho/godotenv"
)

func main() {
	services, err := persistence.ConnUsers()
	if err != nil {
		fmt.Println(err)
	}
	Usersignup := interfaces.NewCountry(services.Country)

	if err := godotenv.Load("config.env"); err != nil {
		panic("Error loading .env file")
	}
	appPort := os.Getenv("port")
	hostAddress := os.Getenv("host_address")
	log.Println("App running on " + hostAddress + appPort)

	appMode := os.Getenv("app_mode")
	if appMode == "dev" {
		r := routes.SetupRouter(appPort, hostAddress, Usersignup)
		http.ListenAndServe(appPort, r)
	}

}
