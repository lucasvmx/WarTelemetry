package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetMapInfoData function retrieves information about current map
func GetMapInfoData(wg *sync.WaitGroup) (mi *mapinfo.MapInformation, err error) {
	defer wg.Done()

	data, err := network.GetDataFromURL(mapinfo.GetURL())
	if err != nil {
		return nil, err
	}

	failure := json.Unmarshal(data, &mi)
	if failure != nil {
		return nil, err
	}

	DataChan <- mi

	return
}
