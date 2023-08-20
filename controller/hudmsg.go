package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/hudmsg"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetHudMessagesData function retrieves all messages from HUD
func GetHudMessagesData() error {
	var messages *hudmsg.Hudmsg

	data, err := network.GetDataFromURL(hudmsg.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get hud message data: %v", err)
	}

	// Decode data into a struct
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return fmt.Errorf("failed to get hud message data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()
	model.TelemetryInstance.HudMessages = messages

	return nil
}
