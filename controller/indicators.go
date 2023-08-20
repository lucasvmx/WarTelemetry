package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/indicators"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData() error {
	var id *indicators.Indicators = &indicators.Indicators{}

	data, err := network.GetDataFromURL(indicators.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get indicators data: %v", err)
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		return fmt.Errorf("failed to get indicators data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.Indicators = id

	return nil
}
