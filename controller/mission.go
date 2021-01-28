package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/mission"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData() (mi *mission.MissionInfo, err error) {
	// Sends a HTTP request
	data, err := network.GetDataFromURL(mission.GetURL())
	if err != nil {
		return nil, err
	}

	// Decode JSON into a struct
	err = json.Unmarshal(data, &mi)
	if err != nil {
		return nil, err
	}

	return
}
