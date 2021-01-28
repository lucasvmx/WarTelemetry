package model

import (
	"github.com/lucas-engen/WarTelemetry/model/gamechat"
	"github.com/lucas-engen/WarTelemetry/model/hudmsg"
	"github.com/lucas-engen/WarTelemetry/model/indicators"
	"github.com/lucas-engen/WarTelemetry/model/mapinfo"
	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
	"github.com/lucas-engen/WarTelemetry/model/mission"
	"github.com/lucas-engen/WarTelemetry/model/state"
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
