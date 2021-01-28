package mission

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

// MissionInfo struct contains data about mission
type MissionInfo struct {
	MissionObjectives []Objectives `json:"objectives"`
	Status            string       `json:"status"`
}

// Objectives struct contains mission objectives
type Objectives struct {
	Primary bool   `json:"primary"`
	Status  string `json:"status"`
	Text    string `json:"text"`
}

var path string = "mission.json"

// GetURL retrieves the URL to get data about mission
func GetURL() string {
	url := utils.GetBaseURL()
	url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
	url = strings.ReplaceAll(url, "$path$", path)
	return url
}
