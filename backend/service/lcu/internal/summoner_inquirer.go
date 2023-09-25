package internal

import (
	"encoding/json"
	"fmt"
	"hidden-record-assistant/backend/model"
)

type IConn interface {
	Get(url string) ([]byte, error)
}

type SummonerInquirer struct {
	IConn
}

func NewSummonerInquirer(conn IConn) *SummonerInquirer {
	return &SummonerInquirer{
		IConn: conn,
	}
}

func (s *SummonerInquirer) GetCurrentSummonerBaseInfo() (data model.SummonerBaseInfo, err error) {
	binData, err := s.Get("/lol-summoner/v1/current-summoner")
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func (s *SummonerInquirer) GetSummonerBaseInfoByName(name string) (data model.SummonerBaseInfo, err error) {
	binData, err := s.Get("/lol-summoner/v1/summoners?name=" + name)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func (s *SummonerInquirer) GetMatchHistory(puuid string, begI, endI int) (data model.MatchHistory, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begI, endI)
	binData, err := s.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func (s *SummonerInquirer) GetMatchDetails(gameId int) (data model.MatchData, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	binData, err := s.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func (s *SummonerInquirer) GetRank(puuid string) (data model.Rank, err error) {
	url := fmt.Sprintf("/lol-ranked/v1/ranked-stats/%s", puuid)
	binData, err := s.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}
