package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/mission"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMissionData function retrieves data about running missions
func GetMissionData(wg *sync.WaitGroup) (mi *mission.MissionInfo, err error) {
	defer wg.Done()

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

	DataChan <- mi

	return
}
