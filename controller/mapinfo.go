package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/mapinfo"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

func GetMapInfoData() (mi *mapinfo.MapInformation) {
	data := network.GetDataFromURL(mapinfo.GetURL())

	failure := json.Unmarshal(data, &mi)
	if failure != nil {
		log.Fatalf("[FATAL] Failed to decode data into a struct: %v", failure)
	}

	return
}
