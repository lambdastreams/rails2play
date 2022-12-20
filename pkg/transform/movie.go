package transform

import (
	"context"
	"encoding/json"
	"fmt"

	b64 "encoding/base64"

	log "github.com/sirupsen/logrus"
	"github.com/subbarao/transformer/pkg/api"
	"github.com/tidwall/gjson"
)

type Movie struct {
	Id              string   `json:"_id"`
	Name            string   `json:"name"`
	Title           string   `json:"title"`
	Rating          string   `json:"rating"`
	Slug            string   `json:"slug"`
	Description     string   `json:"description"`
	Genre           []string `json:"genre"`
	Directors       []string `json:"directors"`
	Writers         []string `json:"writers"`
	Tags            []string `json:"tags"`
	ProgrammingType string   `json:"programming_type"`
}

func contentLookupURL(baseURL string, slug string, contentType string) *api.Builder {
	menu := map[string]any{
		"filter": []map[string]any{
			{"field": "cty", "operator": "equals", "term": contentType},
			{"field": "nu", "operator": "equals", "term": slug},
		},
	}
	data, _ := json.Marshal(&menu)
	query := b64.StdEncoding.EncodeToString([]byte(data))
	url := fmt.Sprintf("%s/content/lookup?reg=us&dt=androidmobile&client=amd-localnow-web&query=%s&pageNumber=1&pageSize=10&sortBy=ut", baseURL, query)

	return api.URL(url)
}

func GetMovieId(baseURL string, slug string) (string, error) {
	contextLogger := log.WithFields(log.Fields{
		"slug": slug,
	})

	var body string
	err := contentLookupURL(baseURL, slug, "movie").ToString(&body).Fetch(context.Background())
	if err != nil {
		contextLogger.Error("Failed to retrieve movie details", err)
		return "", err
	}
	log.WithFields(log.Fields{
		"output": body,
	}).Info("lookup response")
	id := gjson.Get(body, "data.0.id").String()
	log.WithFields(log.Fields{
		"id": id,
	}).Info("resolved movie id")

	return id, nil
}
func GetResource(baseURL string, slug string) (map[string]any, error) {
	movieId, err := GetMovieId(baseURL, slug)
	if err != nil {
		return nil, err
	}
	movie, err := GetMovie(baseURL, movieId)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"channel": []Movie{*movie},
		"success": true,
	}, nil

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
func map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}
func asString(key gjson.Result) string {
	return key.String()
}

func buildMovie(body string) Movie {
	name := gjson.Get(body, "data.lon.#(lang==\"en\").n")
	id := gjson.Get(body, "data.id")
	description := gjson.Get(body, "data.lod.#(lang==\"en\").n")
	rating := gjson.Get(body, "data.rat.0.v")

	genre := map2(gjson.Get(body, "data.log.#(lang==\"en\").n").Array(), asString)
	directors := map2(gjson.Get(body, "data.lodr.#.lon.#(lang==\"en\").n").Array(), asString)
	writers := map2(gjson.Get(body, "data.lowr.#.lon.#(lang==\"en\").n").Array(), asString)
	tags := map2(gjson.Get(body, "data.lotg.#(lang==\"en\").n").Array(), asString)

	slug := gjson.Get(body, "data.nu")
	contentType := gjson.Get(body, "data.cty")

	return Movie{
		Id:              id.String(),
		Name:            name.String(),
		Title:           name.String(),
		Rating:          rating.String(),
		Slug:            slug.String(),
		Genre:           genre,
		Directors:       directors,
		Writers:         writers,
		Tags:            tags,
		ProgrammingType: contentType.String(),
		Description:     description.String(),
	}
}
