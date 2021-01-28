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
	body := client.GetDataFromURL(gamechat.GetURL())

	// Decode data into a struct
	marshalErr := json.Unmarshal(body, &gc)
	if marshalErr != nil {
		log.Fatalf("[FATAL] Failed to marshal JSON data: %v", marshalErr)
	}

	return
}
