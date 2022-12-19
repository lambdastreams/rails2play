package transform

import (
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"testing"

	"github.com/h2non/gock"
)

func TestMoviePropertyTransform(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/movie.json")
	if err != nil {
		t.Fatal(err)
	}
	value := string(b)
	movie := buildMovie(value)
	if movie.Id != "5584D1F9-D627-4205-BDF5-68A541F1BD85" {
		t.Errorf("got %s, wanted %s", movie.Id, "5584D1F9-D627-4205-BDF5-68A541F1BD85")
	}

	if !reflect.DeepEqual(movie.Genre, [3]string{"Horror", "Mystery", "Thriller"}) {
		t.Errorf("got %q, wanted adfk", movie.Genre)
	}

	cases := []struct {
		field    string
		expected string
	}{
		{
			field:    "Id",
			expected: "5584D1F9-D627-4205-BDF5-68A541F1BD85",
		},
		{
			field:    "Name",
			expected: "Tidal Wave (English dub)",
		},
		{
			field:    "Title",
			expected: "Tidal Wave (English dub)",
		},
		{
			field:    "Slug",
			expected: "tidal-wave-english-dub",
		},
		{
			field:    "ProgrammingType",
			expected: "movie",
		},
		{
			field:    "Rating",
			expected: "R",
		},
		{
			field:    "Description",
			expected: "A deep-sea earthquake occurs, creating a tidal wave that is headed straight for Haeundae, a popular vacation spot on the south coast of Korea, which draws visitors from all over the world.",
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

	if movie.Id != "5584D1F9-D627-4205-BDF5-68A541F1BD85" {
		t.Errorf("expected 5584D1F9-D627-4205-BDF5-68A541F1BD85, but got %s", movie.Id)
	}
}
