package model

type InitBackendData struct {
	Auth Auth `json:"auth"`
	LOLJson
}

type Auth struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}

type LOLJson struct {
	ChampionJson map[int]interface{} `json:"champions"`
	QueueJson    map[int]interface{} `json:"queues"`
	PerkJson     map[int]interface{} `json:"perks"`
	ItemJson     map[int]interface{} `json:"items"`
	SpellJson    map[int]interface{} `json:"spells"`
}
