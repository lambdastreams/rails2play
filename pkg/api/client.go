package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ResponseHandler func(*http.Response) error

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
