package main

import (
	"net/http"
	"os"

	"github.com/Keanus-In-Reevesverse/website-structure/routes"
	"github.com/joho/godotenv"
)

//By now loads .env and routes
func main() {
	godotenv.Load(".env")
	routes.LoadRoutes()
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
