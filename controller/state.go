package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/state"
	network "github.com/lucasvmx/WarTelemetry/network/http"
	"github.com/lucasvmx/WarTelemetry/utils"
)

// GetStateData function retrieves data about running missions
func GetStateData(wg *sync.WaitGroup) {
	var st *state.AircraftState
	defer wg.Done()

	// Sends a HTTP request
	data, err := network.GetDataFromURL(state.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get state data: %v", err)
		return
	}

	// Process JSON into a readable format
	data = utils.ProcessJSON(data)

	// Decode JSON into a struct
	err = json.Unmarshal(data, &st)
	if err != nil {
		log.Printf("[ERROR] failed to get state data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.State = st
}
