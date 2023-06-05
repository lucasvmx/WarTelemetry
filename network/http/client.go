package network

import (
	"io"
	"log"
	"net/http"
)

// ret function sends a GET request to the specified URL
func get(url string) (response *http.Response, err error) {
	response, err = http.Get(url)
	if err != nil {
		log.Printf("[ERROR] Failed to send request because of: %v", err)
		return nil, err
	}

	return response, err
}

// readResponse function reads a HTTP request body
func readResponse(response *http.Response) (data []byte, err error) {
	// Release all resources
	defer response.Body.Close()

	// Read response body
	data, err = io.ReadAll(response.Body)

	return
}

// GetDataFromURL function retrieves data from the specified URL as string
func GetDataFromURL(url string) (data []byte, err error) {
	// Sends the GET request
	payload, err := get(url)
	if err != nil {
		return nil, err
	}

	// Reads the HTTP response body
	data, err = readResponse(payload)
	if err != nil {
		return nil, err
	}

	log.Printf("read %d bytes from body (%s)", len(data), url)

	return
}
