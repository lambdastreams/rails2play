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
func Start(port int, baseURL string) {
	r := mux.NewRouter()
	r.HandleFunc("/movie/{id}", getMovieHandler(baseURL))
	r.HandleFunc("/series/{id}", getSeriesHandler(baseURL))
	http.Handle("/", r)
	fmt.Printf("Starting up on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getMovieHandler(baseURL string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]
		movie, err := transform.GetMovie(baseURL, id)
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
}

func getSeriesHandler(baseURL string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params["id"]
		series, err := transform.GetSeries(baseURL, id)
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
}
