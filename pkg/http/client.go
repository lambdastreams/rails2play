package http

import (
	"io/ioutil"
	"net/http"
)

func GetMovie(name string) (string, error) {
	client := &http.Client{}
	url := "https://data-store-cdn.cms-stag.amdvids.com/content/urn/resource/catalog/movie/5584D1F9-D627-4205-BDF5-68A541F1BD85?reg=us&dt=androidmobile&client=amd-localnow-web"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
		return "", err
	}
	req.Header.Add("X-Tracking-Id", "fb8812b9-b5f7-472d-9ab2-8e662253ca03")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
