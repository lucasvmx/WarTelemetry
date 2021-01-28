package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/mapinfo"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetMapInfoData function retrieves information about current map
func GetMapInfoData() (mi *mapinfo.MapInformation, err error) {
	data, err := network.GetDataFromURL(mapinfo.GetURL())
	if err != nil {
		return nil, err
	}

	failure := json.Unmarshal(data, &mi)
	if failure != nil {
		return nil, err
	}

	return
}
