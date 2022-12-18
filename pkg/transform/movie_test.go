package transform

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMoviePropertyTransform(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/test.json")
	if err != nil {
		t.Fatal(err)
	}
	value := string(b)
	movie := buildMovie(value)
	if movie.Id != "5584D1F9-D627-4205-BDF5-68A541F1BD85" {
		t.Errorf("got %q, wanted adfk", movie.Id)
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
