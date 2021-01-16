package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/mission"
	client "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData() (mi *mission.MissionInfo) {
	// Sends a HTTP request
	response := client.Get(mission.GetURL())

	// Reads HTTP response
	data, err := client.ReadResponse(response)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read response because of: %v", err)
	}

	// Decode JSON into a struct
	marshalErr := json.Unmarshal(data, &mi)
	if marshalErr != nil {
		log.Fatalf("[FATAL] Failed to decode JSON data: %v", marshalErr)
	}

	return
}
