package mission

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

// GetURL retrieves the URL to get data about mission
func GetURL() string {
	return "http://localhost:8111/mission.json"
}
