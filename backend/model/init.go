package model

type InitBackendData struct {
	Auth Auth `json:"auth"`
}

type Auth struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}
