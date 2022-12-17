package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetMovie(movie string) (string, error) {
	var result string
	err := URL(fmt.Sprintf("https://data-store-cdn.cms-stag.amdvids.com/content/urn/resource/catalog/movie/%s?reg=us&dt=androidmobile&client=amd-localnow-web", movie)).
		Header("X-Tracking-Id", "fb8812b9-b5f7-472d-9ab2-8e662253ca03").
		ToString(&result).
		Fetch(context.Background())

	if err != nil {
		log.Error("error", err)
		return "", err
	}

	return result, nil

}

type ResponseHandler = func(*http.Response) error

// ToJSON decodes a response as a JSON object.
func ToJSON(v any) ResponseHandler {
	return func(res *http.Response) error {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(data, v); err != nil {
			return err
		}
		return nil
	}
}

// ToString writes the response body to the provided string pointer.
func ToString(sp *string) ResponseHandler {
	return func(res *http.Response) error {
		log.Info("Moving to string")
		var buf strings.Builder
		_, err := io.Copy(&buf, res.Body)
		if err == nil {
			*sp = buf.String()
		} else {
			log.Error("Failed with error", err)
		}
		return err
	}
}
