package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mission"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData(wg *sync.WaitGroup) {
	var mi *mission.MissionInfo
	defer wg.Done()

	// Sends a HTTP request
	data, err := network.GetDataFromURL(mission.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get mission data: %v", err)
		return
	}

	// Decode JSON into a struct
	err = json.Unmarshal(data, &mi)
	if err != nil {
		log.Printf("[ERROR] failed to get mission data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.MissionInfo = mi
}
