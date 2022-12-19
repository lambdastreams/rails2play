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
	app.router.HandleFunc("/movie/{id}", setJSONContentType(app.getMovie))
	app.router.HandleFunc("/series/{id}", setJSONContentType(app.getSeries))
}

// HandleFunc queries quick play for movie details and responds with rails json keys
func (a *App) getMovie(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	movie, err := transform.GetMovie(a.quickPlayURL, id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	var resource = make(map[string]interface{})
	resource["channels"] = [1]transform.Movie{*movie}
	resource["success"] = true

	jsonResponse, jsonError := json.Marshal(resource)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

// HandleFunc queries quick play for series details and responds with rails json keys
func (a *App) getSeries(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	series, err := transform.GetSeries(a.quickPlayURL, id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, jsonError := json.Marshal(series)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func setJSONContentType(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		w.Header().Add("Content-Type", "application/json")
	})
}
