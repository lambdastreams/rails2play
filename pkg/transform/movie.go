package transform

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/subbarao/transformer/pkg/api"
	"github.com/tidwall/gjson"
)

type Movie struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Query movie information from quickplay
func GetMovie(name string) (*Movie, error) {
	contextLogger := log.WithFields(log.Fields{
		"movie": name,
	})

	var body string
	err := api.URL(movieURL(name)).
		ToString(&body).
		Fetch(context.Background())

	if err != nil {
		contextLogger.Error("Failed to retrieve movie details", err)
		return nil, err
	}

	contextLogger.Debug("Building movie details")
	movie := buildMovie(body)

	return &movie, nil
}

func movieURL(movie string) string {
	return fmt.Sprintf("https://data-store-cdn.cms-stag.amdvids.com/content/urn/resource/catalog/movie/%s?reg=us&dt=androidmobile&client=amd-localnow-web", movie)
}

func seriesURL(series string) string {
	return fmt.Sprintf("https://data-store-cdn.cms-stag.amdvids.com/content/series/%s/episodes?reg=us&dt=androidmobile&client=amd-localnow-web&seasonId=00FFFEBA-9E34-4C3E-99F5-D6D814403FD5&pageNumber=1&pageSize=10&sortBy=ut&st=published", series)
}

func buildMovie(body string) Movie {
	name := gjson.Get(body, "data.lon.#(lang==\"en\").n")
	id := gjson.Get(body, "data.id")
	description := gjson.Get(body, "data.lod.#(lang==\"en\").n")

	return Movie{
		Id:          id.String(),
		Name:        name.String(),
		Description: description.String(),
	}
}
