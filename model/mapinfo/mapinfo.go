package mapinfo

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

// MapInformation contains information about map
type MapInformation struct {
	GridSteps []float64 `json:"grid_steps"`
	GridZero  []float64 `json:"grid_zero"`
	MapGen    uint8     `json:"map_generation"`
	MapMax    []float64 `json:"map_max"`
	MapMin    []float64 `json:"map_min"`
	MapSizeX  float64   `json:"map_size_x"`
	MapSizeY  float64   `json:"map_size_y"`

	// MapSize contains the map size (meters)
	MapSize float64 `json:"map_size"`
}

var url string = ""
var imgPath string = "map.img"
var path string = "map_info.json"

// GetURL retrieves URL to get map info data
func GetURL() string {
	url := utils.GetBaseURL()
	url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
	url = strings.ReplaceAll(url, "$path$", path)
	return url
}

// GetImageURL retrieves the image that contains the map drawing
func GetImageURL() string {
	if len(url) == 0 {
		url = utils.GetBaseURL()
		url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
		url = strings.ReplaceAll(url, "$path$", imgPath)
	}

	return url
}

// CalculateMapSize function fills struct with full map size
func (mi *MapInformation) CalculateMapSize() {
	maximumX := mi.MapMax[0]
	minimumY := mi.MapMax[1]
	minimumX := mi.MapMin[0]
	maximumY := mi.MapMin[1]

	// Calculate map sizes
	mi.MapSizeX = maximumX - minimumX
	mi.MapSizeY = maximumY - minimumY
	mi.MapSize = maximumX - minimumX

	if mi.MapSizeX < 0 {
		mi.MapSizeX *= -1.0
	} else if mi.MapSize < 0 {
		mi.MapSize *= -1.0
	} else if mi.MapSizeY < 0 {
		mi.MapSizeY *= -1.0
	}
}
