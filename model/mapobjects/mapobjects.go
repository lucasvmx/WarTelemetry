package mapobjects

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

// GetURL retrieves data about map objects
func GetURL() string {
	return "http://localhost:8111/map_obj.json"
}
