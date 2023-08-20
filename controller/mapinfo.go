package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapInfoData function retrieves information about current map
func GetMapInfoData() error {
	var mi *mapinfo.MapInformation = &mapinfo.MapInformation{}

	data, err := network.GetDataFromURL(mapinfo.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get map information data: %v", err)
	}

	err = json.Unmarshal(data, &mi)
	if err != nil {
		return fmt.Errorf("failed to get map information data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.MapInfo = mi

	return nil
}
