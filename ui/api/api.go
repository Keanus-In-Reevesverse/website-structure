package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetPrices(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(os.Getenv("APIPORT"))

	if err != nil {
		panic(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(responseData)
}
