package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

// ret function sends a GET request to the specified URL
func get(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("[FATAL] Failed to send request because of: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("[INFO] Server returned %d. Url: %v", response.StatusCode, url)
	}

	return response
}

// readResponse function reads a HTTP request body
func readResponse(response *http.Response) (data []byte, err error) {
	// Release all resources
	defer response.Body.Close()

	// Read response body
	data, err = ioutil.ReadAll(response.Body)

	return
}

// GetDataFromURL function retrieves data from the specified URL as string
func GetDataFromURL(url string) []byte {
	payload := get(url)
	data, err := readResponse(payload)
	if err != nil {
		log.Fatalf("[FATAL] Failed to get data from URL: %v", err)
	}

	return data
}
