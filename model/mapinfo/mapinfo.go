package mapinfo

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

type MapInformation struct {
	GridSteps []string `json:"grid_steps"`
	GridZero  []string `json:"grid_zero"`
	MapGen    string   `json:"map_generation"`
	MapMax    []string `json:"map_max"`
	MapMin    []string `json:"map_min"`
}

var path string = "map_info.json"

// Retrieves URL to get map info data
func GetURL() string {
	url := utils.GetBaseURL()
	url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
	url = strings.ReplaceAll(url, "$path$", path)
	return url
}

var img_path string = "map.img"
var url string = ""

// Retrieves the image that contains the map drawing
func GetImageURL() string {
	if len(url) == 0 {
		url = utils.GetBaseURL()
		url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
		url = strings.ReplaceAll(url, "$path$", img_path)
	}

	return url
}
