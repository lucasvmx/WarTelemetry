package indicators

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

// Indicators struct contains data about aircraft indicators
type Indicators struct {
	Valid            bool    `json:"valid"`
	AircraftName     string  `json:"type"`
	Speed            float32 `json:"speed"`
	Pedals           float32 `json:"pedals"`
	Pedals1          float32 `json:"pedals1"`
	Pedals2          float32 `json:"pedals2"`
	Pedals3          float32 `json:"pedals3"`
	StickElevator    float32 `json:"stick_elevator"`
	StickElevator1   float32 `json:"stick_elevator1"`
	StickAilerons    float32 `json:"stick_ailerons"`
	Vario            float32 `json:"vario"`
	AltitudeHour     float32 `json:"altitude_hour"`
	AltitudeMin      float32 `json:"altitude_min"`
	Altitude10k      float32 `json:"altitude_10k"`
	AviaHorizonRoll  float32 `json:"aviahorizon_roll"`
	AviaHorizonPitch float32 `json:"aviahorizon_pitch"`
	Bank             float32 `json:"bank"`
	Turn             float32 `json:"turn"`
	Compass          float32 `json:"compass"`
	Compass1         float32 `json:"compass1"`
	Compass2         float32 `json:"compass2"`
	ClockHour        float32 `json:"clock_hour"`
	ClockMin         float32 `json:"clock_min"`
	ClockSec         float32 `json:"clock_sec"`
	RpmMin           float32 `json:"rpm_min"`
	RpmMin1          float32 `json:"rpm1_min"`
}

var path string = "indicators"

func GetURL() string {
	url := utils.GetBaseURL()
	url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
	url = strings.ReplaceAll(url, "$path$", path)
	return url
}
