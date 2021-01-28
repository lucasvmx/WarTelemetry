package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

var objs *mapobjects.MapObjects

// GetMapObjsData function retrives data about all map objects
func GetMapObjsData() (mo []mapobjects.MapObjects) {

	// Sends GET request
	data := network.GetDataFromURL(mapobjects.GetURL())

	// Decode data into a json struct
	failure := json.Unmarshal(data, &mo)
	if failure != nil {
		log.Fatalf("[FATAL] Failed to decode JSON")
	}

	return
}
