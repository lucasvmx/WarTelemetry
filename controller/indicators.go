package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/indicators"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData(wg *sync.WaitGroup) {
	var id *indicators.Indicators
	defer wg.Done()

	data, err := network.GetDataFromURL(indicators.GetURL())
	if err != nil {
		logger.LogError("failed to get indicators data: %v", err)
		return
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		logger.LogError("failed to get indicators data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.Indicators = id
}
