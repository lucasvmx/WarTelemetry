package gamechat

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

// GameChat struct
type GameChat struct {
	ID      int    `json:"id"`
	Message string `json:"msg"`
	Sender  string `json:"sender"`
	Enemy   bool   `json:"enemy"`
	Mode    string `json:"mode"`
}

// path string
var path string = "gamechat"
var url string = ""

// GetURL function retrieves the URL to get data about gamechat
func GetURL() string {
	if len(url) == 0 {
		url = utils.GetBaseURL()
		url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
		url = strings.ReplaceAll(url, "$path$", path)
	}

	return url
}
