package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subbarao/transformer/pkg/transform"
)

// server to handle rails requests
func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/movie/{id}", handleMovie)
	r.HandleFunc("/series/{id}", handleSeries)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMovie(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	movie, err := transform.GetMovie(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, jsonError := json.Marshal(movie)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func handleSeries(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	series, err := transform.GetSeries(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, jsonError := json.Marshal(series)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
