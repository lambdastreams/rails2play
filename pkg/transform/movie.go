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

func GetMovie(name string) (*Movie, error) {
	var body string
	err := api.URL(fmt.Sprintf("https://data-store-cdn.cms-stag.amdvids.com/content/urn/resource/catalog/movie/%s?reg=us&dt=androidmobile&client=amd-localnow-web", name)).
		ToString(&body).
		Fetch(context.Background())

	if err != nil {
		log.Error("error", err)
		return nil, err
	}
	movie := buildMovie(body)

	return &movie, nil
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
