package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucas-engen/WarTelemetry/model/indicators"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData(wg *sync.WaitGroup) (id *indicators.Indicators, err error) {
	defer wg.Done()

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
