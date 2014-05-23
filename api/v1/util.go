package v1

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Sends a GET request to the query url and returns
// the response or an error.
func sendQuery(query string) ([]byte, error) {
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Returns whether the filepath exists.
func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
