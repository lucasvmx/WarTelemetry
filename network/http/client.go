package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Get function sends a GET request to the specified URL
func Get(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("[FATAL] Failed to send request because of: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("[INFO] Server returned %d. Url: %v", response.StatusCode, url)
	}

	return response
}

// ReadResponse function reads a HTTP request body
func ReadResponse(response *http.Response) (data []byte, err error) {
	// Release all resources
	defer response.Body.Close()

	// Read response body
	data, err = ioutil.ReadAll(response.Body)

	return
}
