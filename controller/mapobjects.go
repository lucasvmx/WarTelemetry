package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData(wg *sync.WaitGroup) {
	var mo []mapobjects.MapObjects
	defer wg.Done()

	// Sends GET request
	data, err := network.GetDataFromURL(mapobjects.GetURL())
	if err != nil {
		logger.LogError("failed to get map objects data: %v", err)
		return
	}

	// Decode data into a json struct
	err = json.Unmarshal(data, &mo)
	if err != nil {
		logger.LogError("failed to get map objects data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.MapObjects = mo
}
