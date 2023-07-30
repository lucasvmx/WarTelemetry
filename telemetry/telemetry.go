package telemetry

import (
	"fmt"
	"sync"
	"time"

	"github.com/lucasvmx/WarTelemetry/controller"
	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/utils"
)

var wg *sync.WaitGroup

var (
	collectors = map[string]func(*sync.WaitGroup){
		"gamechat":   controller.GetGamechatData,
		"indicators": controller.GetIndicatorsData,
		"map_info":   controller.GetMapInfoData,
		"hud_msg":    controller.GetHudMessagesData,
		"map_obj":    controller.GetMapObjsData,
		"state":      controller.GetStateData,
		"mission":    controller.GetMissionData,
	}
)

// Initialize function setup the library to be used
func Initialize(hostname string) {
	if len(hostname) > 0 {
		utils.SetHostname(hostname)
	} else {
		logger.LogInfo("Hostname omitted. Using localhost")
	}

	model.SetupTelemetry()
	wg = &sync.WaitGroup{}
}

func getTelemetryData() {

	start := time.Now()

	for name, callback := range collectors {
		logger.LogInfo("running task go get %v data", name)
		wg.Add(1)
		go callback(wg)
	}

	wg.Wait()

	logger.LogInfo("fetched data in %v", time.Since(start))
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

	if data.MapInfo == nil || data.Gamechat == nil ||
		data.Indicators == nil || data.State == nil ||
		data.HudMessages == nil || data.MapObjects == nil ||
		data.MissionInfo == nil {

		return nil, fmt.Errorf("invalid data returned from some (or all) endpoints")
	}

	return
}
