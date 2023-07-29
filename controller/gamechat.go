package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	client "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData(wg *sync.WaitGroup) {

	var gc []gamechat.GameChat
	defer wg.Done()

	// Sends the GET request
	body, err := client.GetDataFromURL(gamechat.GetURL())
	if err != nil {
		logger.LogError("failed to get gamechat data: %v", err)
		return
	}

	// Decode data into a struct
	err = json.Unmarshal(body, &gc)
	if err != nil {
		logger.LogError("failed to get gamechat data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.Gamechat = gc
}
