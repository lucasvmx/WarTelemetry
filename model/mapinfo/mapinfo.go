package mapinfo

type MapInformation struct {
	GridSteps []string `json:"grid_steps"`
	GridZero  []string `json:"grid_zero"`
	MapGen    string   `json:"map_generation"`
	MapMax    []string `json:"map_max"`
	MapMin    []string `json:"map_min"`
}

// Retrieves URL to get map info data
func GetURL() string {
	return "http://localhost:8111/map_info.json"
}

// Retrieves the image that contains the map drawing
func GetImageURL() string {
	return "http://localhost:8111/map.img"
}
