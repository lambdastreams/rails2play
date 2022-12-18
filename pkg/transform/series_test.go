package transform

import (
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"testing"

	"github.com/h2non/gock"
)

func TestSeriesPropertyTransform(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/series.json")
	if err != nil {
		t.Fatal(err)
	}
	value := string(b)
	series := buildSeries(value)
	cases := []struct {
		field    string
		expected string
	}{
		{
			field:    "Id",
			expected: "1029E3AB-AE97-43BE-A0E9-180D9BA5E688",
		},
		{
			field:    "Name",
			expected: "Tough Day at The Office",
		},
		{
			field:    "Description",
			expected: "Think your day at work was hard? Well check out the nightmare scenarios facing these men and woman as their 9-5 shifts turn into a race to survive. Record-breaking snow, torrential rainfall, and even ",
		},
	}

	for _, tt := range cases {
		t.Run(tt.expected, func(t *testing.T) {
			result := getSeriesField(&series, tt.field)
			if result != tt.expected {
				t.Errorf("%s expected %s, but got %s", tt.field, tt.expected, result)
			}
		})
	}
}

func getSeriesField(v *Series, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
func TestSeries(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://data-store-cdn.cms-stag.amdvids.com/content/series/tedlaslo/episodes?reg=us&dt=androidmobile&client=amd-localnow-web&seasonId=00FFFEBA-9E34-4C3E-99F5-D6D814403FD5&pageNumber=1&pageSize=10&sortBy=ut&st=published").
		Reply(http.StatusOK).
		File(path.Join("testdata", "series.json"))
	movie, _ := GetSeries(baseURL, "tedlaslo")

	if movie.Id != "1029E3AB-AE97-43BE-A0E9-180D9BA5E688" {
		t.Errorf("expected 1029E3AB-AE97-43BE-A0E9-180D9BA5E688, but got %s", movie.Id)
	}
}
