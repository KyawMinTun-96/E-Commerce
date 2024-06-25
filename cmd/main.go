package main

import (
	"ecommerce-app/config"
	"ecommerce-app/migrations"
	"ecommerce-app/models"
	"ecommerce-app/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	config.LoadConfig()                                    // load config
	models.ConnectDatabase(config.GetDBConnectionString()) // connect database

	if len(args) == 2 && args[1] == "migrate" {
		migrations.RunMigrations() // database migrate
	} else if len(args) < 2 {
		router := routes.SetupRouter()
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		fmt.Println("type one arg or none")
	}
}
