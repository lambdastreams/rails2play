package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/h2non/gock"
)

func TestMovieResponseCodes(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://test.quickplay.com/content/urn/resource/catalog/movie/foobar?reg=us&dt=androidmobile&client=amd-localnow-web").
		Reply(http.StatusOK).
		File(path.Join("..", "transform", "testdata", "movie.json"))
	cases := []struct {
		Movie      string
		StatusCode int
	}{
		{
			Movie:      "foobar",
			StatusCode: http.StatusOK,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("movie %s response code: %d", tt.Movie, tt.StatusCode), func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/movie/%s", tt.Movie), nil)
			response := executeRequest(req)
			checkResponseCode(t, response.Code, tt.StatusCode)
			checkResponseContentType(t, response.Header().Get("Content-Type"), "application/json")
		})
	}
}
func TestSeriesResponseCodes(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://test.quickplay.com/content/series/tedlaslo/episodes?reg=us&dt=androidmobile&client=amd-localnow-web&seasonId=00FFFEBA-9E34-4C3E-99F5-D6D814403FD5&pageNumber=1&pageSize=10&sortBy=ut&st=published").
		Reply(http.StatusOK).
		File(path.Join("..", "transform", "testdata", "series.json"))
	cases := []struct {
		Movie      string
		StatusCode int
	}{
		{
			Movie:      "tedlaslo",
			StatusCode: http.StatusOK,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("movie %s response code: %d", tt.Movie, tt.StatusCode), func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/series/%s", tt.Movie), nil)
			response := executeRequest(req)
			checkResponseCode(t, response.Code, tt.StatusCode)
			checkResponseContentType(t, response.Header().Get("Content-Type"), "application/json")
		})
	}
}

var a *App

func TestMain(m *testing.M) {
	a = New("https://test.quickplay.com")

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseContentType(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response code %s. Got %s\n", expected, actual)
	}
}
