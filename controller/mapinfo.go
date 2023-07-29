package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapInfoData function retrieves information about current map
func GetMapInfoData(wg *sync.WaitGroup) {
	var mi *mapinfo.MapInformation = &mapinfo.MapInformation{}
	defer wg.Done()

	data, err := network.GetDataFromURL(mapinfo.GetURL())
	if err != nil {
		logger.LogError("failed to get map information data: %v", err)
		return
	}

	err = json.Unmarshal(data, &mi)
	if err != nil {
		logger.LogError("failed to get map information data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.MapInfo = mi
}
