package telemetry

import (
	"github.com/lucasvmx/WarTelemetry/controller"
	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model"
	"github.com/lucasvmx/WarTelemetry/utils"
	"golang.org/x/sync/errgroup"
)

var errGroup *errgroup.Group

var (
	collectors = map[string]func() error{
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
	errGroup = new(errgroup.Group)
}

func getTelemetryData() error {

	for name, callback := range collectors {
		logger.LogInfo("running task to get %v data", name)
		errGroup.Go(callback)
	}

	return errGroup.Wait()
}

// GetTelemetryData function retrieves all telemetry data
func GetTelemetryData() (data *model.TelemetryData, err error) {

	// Initialize struct
	data = &model.TelemetryData{}

	if fail := getTelemetryData(); fail != nil {
		return nil, fail
	}

	data.MapInfo = model.GetInstance().MapInfo
	data.Gamechat = model.GetInstance().Gamechat
	data.Indicators = model.GetInstance().Indicators
	data.State = model.GetInstance().State
	data.HudMessages = model.GetInstance().HudMessages
	data.MapObjects = model.GetInstance().MapObjects
	data.MissionInfo = model.GetInstance().MissionInfo

	return
}
