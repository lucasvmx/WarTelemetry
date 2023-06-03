package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

var objs *mapobjects.MapObjects

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData(wg *sync.WaitGroup) (mo []mapobjects.MapObjects, err error) {
	defer wg.Done()

	// Sends GET request
	data, err := network.GetDataFromURL(mapobjects.GetURL())
	if err != nil {
		return nil, err
	}

	// Decode data into a json struct
	err = json.Unmarshal(data, &mo)
	if err != nil {
		return nil, err
	}

	DataChan <- mo

	return
}
