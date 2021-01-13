package model

type GameChat struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
	Sender  string `json:"sender"`
	Enemy   bool   `json:"enemy"`
	Mode    string `json:"mode"`
}

// URL contains the network location to retrieve data about game
const URL = "http://localhost:8111/gamechat"
