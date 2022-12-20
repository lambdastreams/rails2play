package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/subbarao/transformer/pkg/transform"
)

// reverse proxy api endpoint
// which constains quickplay related apis
type App struct {
	router       *mux.Router
	quickPlayURL string
}

func New(quickPlayURL string) *App {
	app := App{
		router:       mux.NewRouter(),
		quickPlayURL: quickPlayURL,
	}
	app.initializeRoutes()

	return &app
}

func (a *App) Run(port int) {
	log.WithFields(log.Fields{
		"port":         port,
		"quickPlayURL": a.quickPlayURL,
	}).Info("starting server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.router))
}

// initialize application routes
func (app *App) initializeRoutes() {
	app.router.HandleFunc("/channel/US/{slug}", app.getMovie)
}

// HandleFunc queries quick play for movie details and responds with rails json keys
func (a *App) getMovie(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	slug := params["slug"]
	content, err := transform.GetResource(a.quickPlayURL, slug)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, jsonError := json.Marshal(content)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
