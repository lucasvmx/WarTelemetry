package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/indicators"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData(wg *sync.WaitGroup) {
	var id *indicators.Indicators
	defer wg.Done()

	data, err := network.GetDataFromURL(indicators.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get indicators data: %v", err)
		DataChan <- err
		return
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		log.Printf("[ERROR] failed to get indicators data: %v", err)
		DataChan <- err
		return
	}

	DataChan <- id
}
