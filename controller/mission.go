package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mission"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData() error {
	var mi *mission.MissionInfo

	// Sends a HTTP request
	data, err := network.GetDataFromURL(mission.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get mission data: %v", err)
	}

	// Decode JSON into a struct
	err = json.Unmarshal(data, &mi)
	if err != nil {
		return fmt.Errorf("failed to get mission data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.MissionInfo = mi
	return nil
}
