package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	client "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData(wg *sync.WaitGroup) {

	var gc []gamechat.GameChat = []gamechat.GameChat{}
	defer wg.Done()

	// Sends the GET request
	body, err := client.GetDataFromURL(gamechat.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get gamechat data: %v", err)
		return
	}

	// Decode data into a struct
	json.Unmarshal(body, &gc)

	model.TelemetryInstance.LockMux()
	defer model.TelemetryInstance.UnlockMux()

	model.TelemetryInstance.Gamechat = gc
}
