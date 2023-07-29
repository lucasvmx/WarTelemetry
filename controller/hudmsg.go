package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/hudmsg"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetHudMessagesData function retrieves all messages from HUD
func GetHudMessagesData(wg *sync.WaitGroup) {
	var messages *hudmsg.Hudmsg
	defer wg.Done()

	data, err := network.GetDataFromURL(hudmsg.GetURL())
	if err != nil {
		logger.LogError("failed to get hud message data: %v", err)
		return
	}

	// Decode data into a struct
	err = json.Unmarshal(data, &messages)
	if err != nil {
		logger.LogError("failed to get hud message data: %v", err)
		return
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.HudMessages = messages
}
