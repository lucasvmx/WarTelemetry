package mapobjects

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

// MapObjects struct contains data about map objects
type MapObjects struct {
	Type   string  `json:"type"`
	Color  string  `json:"color"`
	Colors []int   `json:"color[]"`
	Blink  int     `json:"blink"`
	Icon   string  `json:"icon"`
	IconBG string  `json:"icon_bg"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Dx     float32 `json:"dx"`
	Dy     float32 `json:"dy"`
	Sx     float32 `json:"sx"`
	Sy     float32 `json:"sy"`
	Ex     float32 `json:"ex"`
	Ey     float32 `json:"ey"`
}

var path string = "map_obj.json"
var url string = ""

// GetURL retrieves data about map objects
func GetURL() string {
	if len(url) == 0 {
		url = utils.GetBaseURL()
		url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
		url = strings.ReplaceAll(url, "$path$", path)
	}

	return url
}
