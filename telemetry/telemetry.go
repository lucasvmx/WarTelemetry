package telemetry

import (
	"log"
	"sync"

	"github.com/lucasvmx/WarTelemetry/controller"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	"github.com/lucasvmx/WarTelemetry/model/hudmsg"
	"github.com/lucasvmx/WarTelemetry/model/indicators"
	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	"github.com/lucasvmx/WarTelemetry/model/mission"
	"github.com/lucasvmx/WarTelemetry/model/state"
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

	wg = &sync.WaitGroup{}
}

// GetTelemetryData function retrieves all telemetry data
func GetTelemetryData() (data *model.TelemetryData, err error) {

	wg.Add(7)

	go controller.GetGamechatData(wg)
	go controller.GetIndicatorsData(wg)
	go controller.GetMapInfoData(wg)
	go controller.GetHudMessagesData(wg)
	go controller.GetMapObjsData(wg)
	go controller.GetStateData(wg)
	go controller.GetMissionData(wg)

	wg.Wait()

	// Initialize struct
	data = &model.TelemetryData{}

	for e := range controller.DataChan {
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
	}

	return
}
