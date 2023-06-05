package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapInfoData function retrieves information about current map
func GetMapInfoData(wg *sync.WaitGroup) {
	var mi *mapinfo.MapInformation
	defer wg.Done()

	data, err := network.GetDataFromURL(mapinfo.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get map information data: %v", err)
		DataChan <- err
		return
	}

	err = json.Unmarshal(data, &mi)
	if err != nil {
		log.Printf("[ERROR] failed to get map information data: %v", err)
		DataChan <- err
		return
	}

	DataChan <- mi
}
