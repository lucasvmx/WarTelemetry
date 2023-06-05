package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData(wg *sync.WaitGroup) {
	var mo []mapobjects.MapObjects = []mapobjects.MapObjects{}
	defer wg.Done()

	// Sends GET request
	data, err := network.GetDataFromURL(mapobjects.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get map objects data: %v", err)
		return
	}

	// Decode data into a json struct
	json.Unmarshal(data, &mo)

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.MapObjects = mo
}
