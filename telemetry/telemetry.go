package telemetry

import (
	"log"

	"github.com/lucas-engen/WarTelemetry/controller"
	"github.com/lucas-engen/WarTelemetry/model"
	"github.com/lucas-engen/WarTelemetry/model/gamechat"
	"github.com/lucas-engen/WarTelemetry/model/hudmsg"
	"github.com/lucas-engen/WarTelemetry/model/indicators"
	"github.com/lucas-engen/WarTelemetry/model/mapinfo"
	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
	"github.com/lucas-engen/WarTelemetry/model/mission"
	"github.com/lucas-engen/WarTelemetry/model/state"
	"github.com/lucas-engen/WarTelemetry/utils"
)

// Initialize function setup the library to be used
func Initialize(hostname string) {
	if len(hostname) > 0 {
		utils.SetHostname(hostname)
	} else {
		log.Printf("[INFO] Hostname omitted. Using localhost")
	}
}

// GetTelemetryData function retrieves all telemetry data
func GetTelemetryData() (data *model.TelemetryData, err error) {
	var count int = 0
	go controller.GetGamechatData()
	go controller.GetIndicatorsData()
	go controller.GetMapInfoData()
	go controller.GetHudMessagesData()
	go controller.GetMapObjsData()
	go controller.GetStateData()
	go controller.GetMissionData()

	// Initialize struct
	data = &model.TelemetryData{}

	for count < 7 {

		e := <-controller.DataChan

		switch value := e.(type) {
		case *mapinfo.MapInformation:
			data.MapInfo = value
		case []gamechat.GameChat:
			data.Gamechat = value
		case *indicators.Indicators:
			data.Indicators = value
		case *state.AircraftState:
			data.State = value
		case *hudmsg.Hudmsg:
			data.HudMessages = value
		case []mapobjects.MapObjects:
			data.MapObjects = value
		case *mission.MissionInfo:
			data.MissionInfo = value
		}

		count++
	}

	return
}
