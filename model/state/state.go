package state

// AircraftState struct contains data about current vehicle state
type AircraftState struct {
	Valid             bool    `json:"valid"`
	AileronPercent    float32 `json:"aileron"`
	ElevatorPercent   float32 `json:"elevator"`
	RudderPercent     float32 `json:"rudder"`
	FlapsPercent      float32 `json:"flaps"`
	Altitude          float32 `json:"H"`
	TrueAirspeed      uint32  `json:"TAS"`
	IndicatedAirspeed uint32  `json:"IAS"`
	Momentum          float32 `json:"M"`
	AngleOfAttack     float32 `json:"AoA"`
	AngleOfSlip       float32 `json:"AoS"`
	Ny                float32 `json:"Ny"`
	VerticalSpeed     float32 `json:"Vy"`
	AngularSpeedX     float32 `json:"Wx"`
	Fuel              uint32  `json:"Mfuel"`
	Fuel0             float32 `json:"Mfuel0"`

	// Engine 1
	ThrottleEngine1   float32 `json:"throttle1"`
	MixtureEngine1    float32 `json:"mixture1"`
	Magneto1          float32 `json:"magneto1"`
	PowerEngine1      float32 `json:"power1"`
	RPMEngine1        float32 `json:"RPM1"`
	ManifoldPressure1 float32 `json:"manifoldpressure1"`
	OilTemp1          float32 `json:"oiltemp1"`
	Pitch1            float32 `json:"pitch1"`
	Thrust1           float32 `json:"thrust1"`
	Efficiency1       float32 `json:"efficiency1"`

	// Engine 2
	ThrottleEngine2   float32 `json:"throttle2"`
	MixtureEngine2    float32 `json:"mixture2"`
	Magneto2          float32 `json:"magneto2"`
	PowerEngine2      float32 `json:"power2"`
	RPMEngine2        float32 `json:"RPM2"`
	ManifoldPressure2 float32 `json:"manifoldpressure2"`
	OilTemp2          float32 `json:"oiltemp2"`
	Pitch2            float32 `json:"pitch2"`
	Thrust2           float32 `json:"thrust2"`
	Efficiency2       float32 `json:"efficiency2"`

	// Engine 3
	ThrottleEngine3   float32 `json:"throttle3"`
	MixtureEngine3    float32 `json:"mixture3"`
	Magneto3          float32 `json:"magneto3"`
	PowerEngine3      float32 `json:"power3"`
	RPMEngine3        float32 `json:"RPM3"`
	ManifoldPressure3 float32 `json:"manifoldpressure3"`
	OilTemp3          float32 `json:"oiltemp3"`
	Pitch3            float32 `json:"pitch3"`
	Thrust3           float32 `json:"thrust3"`
	Efficiency3       float32 `json:"efficiency3"`

	// Engine 4
	ThrottleEngine4   float32 `json:"throttle4"`
	MixtureEngine4    float32 `json:"mixture4"`
	Magneto4          float32 `json:"magneto4"`
	PowerEngine4      float32 `json:"power4"`
	RPMEngine4        float32 `json:"RPM4"`
	ManifoldPressure4 float32 `json:"manifoldpressure4"`
	OilTemp4          float32 `json:"oiltemp4"`
	Pitch4            float32 `json:"pitch4"`
	Thrust4           float32 `json:"thrust4"`
	Efficiency4       float32 `json:"efficiency4"`
}

// GetURL returns the full URL path to access vehicle state resources
func GetURL() string {
	return "http://localhost:8111/state"
}
