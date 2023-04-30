package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/indicators"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData() (id *indicators.Indicators, err error) {
	data, err := network.GetDataFromURL(indicators.GetURL())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		return nil, err
	}

	DataChan <- id

	return
}
