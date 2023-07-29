package network

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lucasvmx/WarTelemetry/logger"
)

// ret function sends a GET request to the specified URL
func get(url string) (response *http.Response, err error) {
	response, err = http.Get(url)
	if err != nil {
		logger.LogError("Failed to send request because of: %v", err)
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

	if payload.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received unexpected HTTP code: %v", payload.StatusCode)
	}

	// Reads the HTTP response body
	data, err = readResponse(payload)
	if err != nil {
		return nil, err
	}

	return
}
