package wartelemetry

import (
	"log"

	"github.com/lucas-engen/WarTelemetry/controller"
	"github.com/lucas-engen/WarTelemetry/model"
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
func GetTelemetryData() (data *model.TelemetryData) {
	gamechat := controller.GetGamechatData()
	indicators := controller.GetIndicatorsData()
	mapInfo := controller.GetMapInfoData()
	hudMessages := controller.GetHudMessagesData()
	mapObjects := controller.GetMapObjsData()
	state := controller.GetStateData()
	missionData := controller.GetMissionData()

	data.MapInfo = mapInfo
	data.Gamechat = gamechat
	data.Indicators = indicators
	data.HudMessages = hudMessages
	data.State = state
	data.MapObjects = mapObjects
	data.MissionInfo = missionData

	return
}
