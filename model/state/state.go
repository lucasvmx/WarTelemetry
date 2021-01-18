package state

// AircraftState struct contains data about current vehicle state
type AircraftState struct {
	Valid             bool    `json:"valid"`
	AileronPercent    int     `json:"aileron, %"`
	ElevatorPercent   int     `json:"elevator, %"`
	RudderPercent     int     `json:"rudder, %"`
	Altitude          int     `json:"H, m"`
	TrueAirspeed      uint32  `json:"TAS, km/h"`
	IndicatedAirspeed uint32  `json:"IAS, km/h"`
	Momentum          int     `json:"M"`
	AngleOfAttack     float32 `json:"AoA, deg"`
	AngleOfSlip       float32 `json:"AoS, deg"`
	Ny                float32 `json:"Ny"`
	VerticalSpeed     float32 `json:"Vy, m/s"`
	AngularSpeedX     float32 `json:"Wx, deg/s"`
	Fuel              uint32  `json:"Mfuel, kg"`
	Fuel0             float32 `json:"Mfuel0, kg"`
	ThrottleEngine1   int     `json:"throttle 1, %"`
	ThrottleEngine2   int     `json:"throttle 2, %"`
	RadiatorEngine1   int     `json:"radiator 1, %"`
	RadiatorEngine2   int     `json:"radiator 2, %"`
	Magneto           int     `json:"magneto 1"`
	PowerEngine1      float32 `json:"power 1, hp"`
	PowerEngine2      float32 `json:"power 2, hp"`
	RpmEngine1        int     `json:"RPM 1"`
	RpmEngine2        int     `json:"RPM 2"`
	ManifoldPressure1 float32 `json:"manifold pressure 1, atm"`
	ManifoldPressure2 float32 `json:"manifold pressure 2, atm"`
	WaterTemp1        float32 `json:"water temp 1, C"`
	WaterTemp2        float32 `json:"water temp 2, C"`
	OilTemp1          int     `json:"oil temp 1, C"`
	Pitch             float32 `json:"pitch 1, deg"`
	Thrust1           int     `json:"thrust 1, kgs"`
	Thrust2           int     `json:"thrust 2, kgs"`
	Efficiency1       int     `json:"efficiency 1, %"`
	Efficiency2       int     `json:"efficiency 2, %"`
}

// GetURL returns the full URL path to access vehicle state resources
func GetURL() string {
	return "http://localhost:8111/state"
}
