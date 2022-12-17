package rails

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type Movie struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (b Movie) String() string {
	return fmt.Sprintf("%b", b)
}

func BuildMovie(body string) Movie {
	name := gjson.Get(body, "data.lon.#(lang==\"en\").n")
	id := gjson.Get(body, "data.id")
	description := gjson.Get(body, "data.lod.#(lang==\"en\").n")
	println(id.String())

	return Movie{
		Id:          id.String(),
		Name:        name.String(),
		Description: description.String(),
	}
}
