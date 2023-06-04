package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucas-engen/WarTelemetry/model/state"
	network "github.com/lucas-engen/WarTelemetry/network/http"
	"github.com/lucas-engen/WarTelemetry/utils"
)

// GetStateData function retrieves data about running missions
func GetStateData(wg *sync.WaitGroup) (st *state.AircraftState, err error) {
	defer wg.Done()

	// Sends a HTTP request
	data, err := network.GetDataFromURL(state.GetURL())
	if err != nil {
		return nil, err
	}

	// Process JSON into a readable format
	data = utils.ProcessJSON(data)

	// Decode JSON into a struct
	err = json.Unmarshal(data, &st)
	if err != nil {
		return nil, err
	}

	return
}
