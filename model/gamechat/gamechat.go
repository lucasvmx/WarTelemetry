package gamechat

// GameChat struct
type GameChat struct {
	ID      int    `json:"id"`
	Message string `json:"msg"`
	Sender  string `json:"sender"`
	Enemy   bool   `json:"enemy"`
	Mode    string `json:"mode"`
}

// GetURL function retrieves the URL to get data about gamechat
func GetURL() string {
	return "http://localhost:8111/gamechat"
}
