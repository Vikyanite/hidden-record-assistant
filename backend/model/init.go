package model

type InitBackendData struct {
	Auth            Auth     `json:"auth"`
	CurrentSummoner Summoner `json:"currentSummoner"`
}

type Auth struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}
