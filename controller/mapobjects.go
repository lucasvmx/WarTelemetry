package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

var objs *mapobjects.MapObjects

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData() (mo []mapobjects.MapObjects, err error) {

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
