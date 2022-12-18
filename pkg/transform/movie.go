package transform

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/subbarao/transformer/pkg/api"
	"github.com/tidwall/gjson"
)

type Movie struct {
	Id              string `json:"_id"`
	Name            string `json:"name"`
	Title           string `json:"title"`
	Rating          string `json:"rating"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	ProgrammingType string `json:"programming_type"`
}

// Query movie information from quickplay
func GetMovie(baseURL string, name string) (*Movie, error) {
	contextLogger := log.WithFields(log.Fields{
		"movie": name,
	})

	var body string
	err := api.URL(movieURL(baseURL, name)).
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

func movieURL(baseURL string, movie string) string {
	return fmt.Sprintf("%s/content/urn/resource/catalog/movie/%s?reg=us&dt=androidmobile&client=amd-localnow-web", baseURL, movie)
}

func seriesURL(baseURL string, series string) string {
	return fmt.Sprintf("%s/content/series/%s/episodes?reg=us&dt=androidmobile&client=amd-localnow-web&seasonId=00FFFEBA-9E34-4C3E-99F5-D6D814403FD5&pageNumber=1&pageSize=10&sortBy=ut&st=published", baseURL, series)
}

func buildMovie(body string) Movie {
	name := gjson.Get(body, "data.lon.#(lang==\"en\").n")
	id := gjson.Get(body, "data.id")
	description := gjson.Get(body, "data.lod.#(lang==\"en\").n")
	rating := gjson.Get(body, "data.rat.0.v")
	slug := gjson.Get(body, "data.nu")
	contentType := gjson.Get(body, "data.cty")

	return Movie{
		Id:              id.String(),
		Name:            name.String(),
		Title:           name.String(),
		Rating:          rating.String(),
		Slug:            slug.String(),
		ProgrammingType: contentType.String(),
		Description:     description.String(),
	}
}
