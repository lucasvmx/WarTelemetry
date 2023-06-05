package telemetry

import (
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/controller"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/utils"
)

var wg *sync.WaitGroup

// Initialize function setup the library to be used
func Initialize(hostname string) {
	if len(hostname) > 0 {
		utils.SetHostname(hostname)
	} else {
		log.Printf("[INFO] Hostname omitted. Using localhost")
	}

	model.SetupTelemetry()
	wg = &sync.WaitGroup{}
}

func getTelemetryData() {
	wg.Add(7)
	go controller.GetGamechatData(wg)
	go controller.GetIndicatorsData(wg)
	go controller.GetMapInfoData(wg)
	go controller.GetHudMessagesData(wg)
	go controller.GetMapObjsData(wg)
	go controller.GetStateData(wg)
	go controller.GetMissionData(wg)
	wg.Wait()
}

// GetTelemetryData function retrieves all telemetry data
func GetTelemetryData() (data *model.TelemetryData, err error) {

	// Initialize struct
	data = &model.TelemetryData{}

	getTelemetryData()

	data.MapInfo = model.GetInstance().MapInfo
	data.Gamechat = model.GetInstance().Gamechat
	data.Indicators = model.GetInstance().Indicators
	data.State = model.GetInstance().State
	data.HudMessages = model.GetInstance().HudMessages
	data.MapObjects = model.GetInstance().MapObjects
	data.MissionInfo = model.GetInstance().MissionInfo

	return
}
