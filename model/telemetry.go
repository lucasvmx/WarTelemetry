package model

import (
	"github.com/lucasvmx/WarTelemetry/model/gamechat"
	"github.com/lucasvmx/WarTelemetry/model/hudmsg"
	"github.com/lucasvmx/WarTelemetry/model/indicators"
	"github.com/lucasvmx/WarTelemetry/model/mapinfo"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	"github.com/lucasvmx/WarTelemetry/model/mission"
	"github.com/lucasvmx/WarTelemetry/model/state"
)

// TelemetryData struct contains all telemetry information
type TelemetryData struct {
	Gamechat    []gamechat.GameChat
	MapInfo     *mapinfo.MapInformation
	Indicators  *indicators.Indicators
	HudMessages *hudmsg.Hudmsg
	State       *state.AircraftState
	MapObjects  []mapobjects.MapObjects
	MissionInfo *mission.MissionInfo
}
