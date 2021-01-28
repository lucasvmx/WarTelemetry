package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/hudmsg"
	network "github.com/lucas-engen/WarTelemetry/network/http"
)

func GetHudMessagesData() (messages *hudmsg.Hudmsg) {

	data := network.GetDataFromURL(hudmsg.GetURL())

	// Decode data into a struct
	failure := json.Unmarshal(data, &messages)
	if failure != nil {
		log.Fatalf("[FATAL] Failed to retrive data from server: %v", failure)
	}

	return
}
