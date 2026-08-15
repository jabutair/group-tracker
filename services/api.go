package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


// FetchJSON gets data from API and decodes JSON
func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, target)
}
