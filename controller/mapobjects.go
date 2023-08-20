package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData() error {
	var mo []mapobjects.MapObjects = []mapobjects.MapObjects{}

	// Sends GET request
	data, err := network.GetDataFromURL(mapobjects.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get map objects data: %v", err)
	}

	// Decode data into a json struct
	err = json.Unmarshal(data, &mo)
	if err != nil {
		return fmt.Errorf("failed to get map objects data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.MapObjects = mo

	return nil
}
