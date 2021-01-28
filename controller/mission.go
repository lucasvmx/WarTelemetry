package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/mission"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData() (mi *mission.MissionInfo) {
	// Sends a HTTP request
	data := network.GetDataFromURL(mission.GetURL())

	// Decode JSON into a struct
	marshalErr := json.Unmarshal(data, &mi)
	if marshalErr != nil {
		log.Fatalf("[FATAL] Failed to decode JSON data: %v", marshalErr)
	}

	return
}
