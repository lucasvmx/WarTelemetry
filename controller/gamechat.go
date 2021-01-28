package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/gamechat"
	client "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData() (gc []gamechat.GameChat) {
	// Sends the GET request
	response := client.Get(gamechat.GetURL())

	log.Printf("[INFO] Response size: %d", response.ContentLength)

	// Read response body
	body, err := client.ReadResponse(response)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read HTTP response because of: %v", err)
	}

	// Decode data into a struct
	marshalErr := json.Unmarshal(body, &gc)
	if marshalErr != nil {
		log.Fatalf("[FATAL] Failed to marshal JSON data: %v", marshalErr)
	}

	return
}
