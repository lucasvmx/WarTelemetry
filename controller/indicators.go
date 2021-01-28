package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/indicators"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetIndicatorsData function retrieves data about aircraft indicators
func GetIndicatorsData() (id *indicators.Indicators) {
	data := network.GetDataFromURL(indicators.GetURL())

	err := json.Unmarshal(data, &id)
	if err != nil {
		log.Fatalf("[FATAL] Failed to decode data into a struct: %v", err)
	}

	return
}
