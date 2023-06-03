package controller

import (
	"encoding/json"
	"sync"

	"github.com/lucasvmx/WarTelemetry/model/hudmsg"
	network "github.com/lucasvmx/WarTelemetry/network/http"
)

// GetHudMessagesData function retrieves all messages from HUD
func GetHudMessagesData(wg *sync.WaitGroup) (messages *hudmsg.Hudmsg, err error) {

	defer wg.Done()

	data, err := network.GetDataFromURL(hudmsg.GetURL())
	if err != nil {
		return nil, err
	}

	// Decode data into a struct
	failure := json.Unmarshal(data, &messages)
	if failure != nil {
		return nil, err
	}

	DataChan <- messages

	return
}
