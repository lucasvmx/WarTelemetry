package controller

import (
	"encoding/json"

	"github.com/lucas-engen/WarTelemetry/model/gamechat"
	client "github.com/lucas-engen/WarTelemetry/network/http"
)

// GetGamechatData retrieves data about gamechat
func GetGamechatData() (gc []gamechat.GameChat, err error) {
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

	return
}
