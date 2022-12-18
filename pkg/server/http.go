package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subbarao/transformer/pkg/api"
	"github.com/subbarao/transformer/pkg/transform"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/movie/{id}", handleMovie)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleMovie(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	movie, err := api.GetMovie(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, jsonError := json.Marshal(transform.BuildMovie(movie))

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
