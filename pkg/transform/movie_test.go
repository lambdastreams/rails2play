package transform

import (
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h2non/gock"
)

type Filter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Term     string `json:"term"`
}

func TestQuerySlugName(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://quickplay.com/content/lookup?reg=us&dt=androidmobile&client=amd-localnow-web&query=eyJmaWx0ZXIiOlt7ImZpZWxkIjoiY3R5Iiwib3BlcmF0b3IiOiJlcXVhbHMiLCJ0ZXJtIjoibW92aWUifSx7ImZpZWxkIjoibnUiLCJvcGVyYXRvciI6ImVxdWFscyIsInRlcm0iOiJzdWJsZXQtdGhlIn1dfQ==&pageNumber=1&pageSize=10&sortBy=ut").
		Reply(http.StatusOK).
		File(path.Join("testdata", "movie_lookup.json"))
	result, _ := GetMovieId("https://quickplay.com", "sublet-the")

	assert.Equal(t, result, "")
}
func TestMoviePropertyTransform(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/movie.json")
	if err != nil {
		t.Fatal(err)
	}
	value := string(b)
	movie := buildMovie(value)

	assert.Equal(t, movie.Genre, []string{"Horror", "Mystery", "Thriller"})
	cases := []struct {
		field    string
		expected string
	}{
		{
			field:    "Id",
			expected: "EF392B33-F6AB-4323-8DF9-3F9F761FFFD4",
		},
		{
			field:    "Name",
			expected: "The Sublet",
		},
		{
			field:    "Title",
			expected: "The Sublet",
		},
		{
			field:    "Slug",
			expected: "sublet-the",
		},
		{
			field:    "ProgrammingType",
			expected: "movie",
		},
		{
			field:    "Rating",
			expected: "TV-14",
		},
		{
			field:    "Description",
			expected: "The Sublet is a suspense driven psychological thriller about Joanna, a new mom coping with her baby alone in an odd sublet apartment. As her husband neglects her to focus on his career, Joanna questio",
		},
	}

	for _, tt := range cases {
		t.Run(tt.expected, func(t *testing.T) {
			result := getField(&movie, tt.field)
			if result != tt.expected {
				t.Errorf("%s expected %s, but got %s", tt.field, tt.expected, result)
			}
		})
	}
}

func getField(v *Movie, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

const baseURL = "https://data-store-cdn.cms-stag.amdvids.com/"

func TestMovie(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://data-store-cdn.cms-stag.amdvids.com/content/urn/resource/catalog/movie/foobar?reg=us&dt=androidmobile&client=amd-localnow-web").
		Reply(http.StatusOK).
		File(path.Join("testdata", "movie.json"))
	movie, _ := GetMovie(baseURL, "foobar")
	assert.Equal(t, movie.Id, "EF392B33-F6AB-4323-8DF9-3F9F761FFFD4")
}
