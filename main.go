package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	ourhttp "github.com/subbarao/transformer/pkg/http"
	"github.com/subbarao/transformer/pkg/rails"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movie/{id}", sendResponse)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendResponse(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	movieId := params["id"]
	movie, _ := ourhttp.GetMovie(movieId)

	jsonResponse, jsonError := json.Marshal(rails.BuildMovie(movie))

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
