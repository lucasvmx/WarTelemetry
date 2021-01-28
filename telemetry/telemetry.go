package wartelemetry

import (
	"github.com/lucas-engen/WarTelemetry/controller"
	"github.com/lucas-engen/WarTelemetry/model"
)

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
