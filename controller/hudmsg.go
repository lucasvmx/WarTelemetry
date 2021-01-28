package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/hudmsg"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetHudMessagesData function retrieves all messages from HUD
func GetHudMessagesData() (messages *hudmsg.Hudmsg, err error) {

	data, err := network.GetDataFromURL(hudmsg.GetURL())
	if err != nil {
		return nil, err
	}

	// Decode data into a struct
	failure := json.Unmarshal(data, &messages)
	if failure != nil {
		return nil, err
	}

	return
}
