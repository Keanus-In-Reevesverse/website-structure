package routes

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/ui/api"
)

func LoadRoutes() {
	//Prices API
	http.HandleFunc("/", api.GetPrices)
	http.HandleFunc("/steam", nil)
	http.HandleFunc("/epic", nil)
	http.HandleFunc("/psn", nil)
	http.HandleFunc("/xbox", nil)
	http.HandleFunc("/nuuvem", nil)

	//Page ammount
	http.HandleFunc("/games", nil)
	http.HandleFunc("/store/steam", nil)
	http.HandleFunc("/store/epic", nil)
	http.HandleFunc("/store/psn", nil)
	http.HandleFunc("/store/xbox", nil)
	http.HandleFunc("/store/nuuvem", nil)

	//Utilities
	http.HandleFunc("/users", nil)
	http.HandleFunc("/users/{id}", nil)
}
