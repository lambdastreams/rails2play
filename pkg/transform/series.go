package transform

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/subbarao/transformer/pkg/api"
	"github.com/tidwall/gjson"
)

type Series struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetSeries(baseURL string, name string) (*Series, error) {
	var body string
	err := api.URL(seriesURL(baseURL, name)).
		ToString(&body).
		Fetch(context.Background())

	if err != nil {
		log.Error("error", err)
		return nil, err
	}
	series := buildSeries(body)

	return &series, nil
}

func buildSeries(body string) Series {
	name := gjson.Get(body, "data.0.lon.#(lang==\"en\").n")
	id := gjson.Get(body, "data.0.id")
	description := gjson.Get(body, "data.0.lod.#(lang==\"en\").n")

	return Series{
		Id:          id.String(),
		Name:        name.String(),
		Description: description.String(),
	}
}
