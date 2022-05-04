package main

import (
	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/routes"
	"github.com/joho/godotenv"
)

//Database connect and handlers
func main() {
	godotenv.Load(".env")
	database.DatabaseConnect()
	routes.HandleRequests()
}
