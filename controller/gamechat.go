package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	client "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData(wg *sync.WaitGroup) (gc []gamechat.GameChat, err error) {

	defer wg.Done()

	// Sends the GET request
	body, err := client.GetDataFromURL(gamechat.GetURL())
	if err != nil {
		return nil, err
	}

	// Decode data into a struct
	err = json.Unmarshal(body, &gc)
	if err != nil {
		return nil, err
	}

	DataChan <- gc

	return
}
