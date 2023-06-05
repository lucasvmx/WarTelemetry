package controller

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	client "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData(wg *sync.WaitGroup) {

	var gc []gamechat.GameChat
	defer func() {
		log.Printf("cleaning up context")
		wg.Done()
	}()

	// Sends the GET request
	body, err := client.GetDataFromURL(gamechat.GetURL())
	if err != nil {
		log.Printf("[ERROR] failed to get gamechat data: %v", err)
		DataChan <- err
		return
	}

	// Decode data into a struct
	err = json.Unmarshal(body, &gc)
	if err != nil {
		log.Printf("[ERROR] failed to get gamechat data: %v", err)
		DataChan <- err
		return
	}

	DataChan <- gc
}
