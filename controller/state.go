package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/state"
	network "github.com/lucasvmx/WarTelemetry/network/http"
	"github.com/lucasvmx/WarTelemetry/utils"
)

// GetStateData function retrieves data about running missions
func GetStateData() error {
	var st *state.AircraftState = &state.AircraftState{}

	// Sends a HTTP request
	data, err := network.GetDataFromURL(state.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get state data: %v", err)
	}

	// Process JSON into a readable format
	data = utils.ProcessJSON(data)

	// Decode JSON into a struct
	err = json.Unmarshal(data, &st)
	if err != nil {
		return fmt.Errorf("failed to get state data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.State = st

	return nil
}
