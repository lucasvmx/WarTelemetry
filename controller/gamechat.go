package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	client "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData() error {

	var gc []gamechat.GameChat = []gamechat.GameChat{}

	// Sends the GET request
	body, err := client.GetDataFromURL(gamechat.GetURL())
	if err != nil {
		return fmt.Errorf("failed to get gamechat data: %v", err)
	}

	// Decode data into a struct
	err = json.Unmarshal(body, &gc)
	if err != nil {
		return fmt.Errorf("failed to get gamechat data: %v", err)
	}

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.Gamechat = gc

	return nil
}
