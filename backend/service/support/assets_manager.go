package support

import (
	"encoding/json"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/module/errs"
	"strconv"
)

type AssetsManager struct {
	model.Assets
	c *Connector
}

func NewAssetsManager(c *Connector) (m *AssetsManager) {
	m = &AssetsManager{
		c: c,
		Assets: model.Assets{
			Champions: map[int]model.Champion{},
			Queues:    map[int]model.Queue{},
			Perks:     map[int]model.Perk{},
			Items:     map[int]model.Item{},
			Spells:    map[int]model.Spell{},
			PerkStyle: map[int]model.PerkStyle{
				8000: {
					Description: "精密",
					IconPath:    "/lol-game-data/assets/v1/perk-images/Styles/7201_precision.png",
				},
				8100: {
					Description: "主宰",
					IconPath:    "/lol-game-data/assets/v1/perk-images/Styles/7200_domination.png",
				},
				8200: {
					Description: "巫术",
					IconPath:    "/lol-game-data/assets/v1/perk-images/Styles/7202_sorcery.png",
				},
				8300: {
					Description: "启迪",
					IconPath:    "/lol-game-data/assets/v1/perk-images/Styles/7203_whimsy.png",
				},
				8400: {
					Description: "坚决",
					IconPath:    "/lol-game-data/assets/v1/perk-images/Styles/7204_resolve.png",
				},
			},
		},
	}
	return
}

func (m *AssetsManager) Init() error {
	errChan, num := ExecuteTaskConcurrently(
		func() (err error) {
			binData, err := m.c.Get("/lol-game-data/assets/v1/champion-summary.json")
			if err != nil {
				return errs.UnknownError(err)
			}
			array := []model.Champion{}
			err = json.Unmarshal(binData, &array)
			if err != nil {
				return errs.UnknownError(err)
			}
			for i := range array {
				m.Assets.Champions[array[i].Id] = array[i]
			}
			return
		},
		func() (err error) {
			binData, err := m.c.Get("/lol-game-data/assets/v1/queues.json")
			if err != nil {
				return errs.UnknownError(err)
			}
			mp := map[string]model.Queue{}
			err = json.Unmarshal(binData, &mp)
			if err != nil {
				return errs.UnknownError(err)
			}
			for k, v := range mp {
				id, _ := strconv.Atoi(k)
				v.Id = id
				m.Assets.Queues[id] = v
			}
			return
		},
		func() (err error) {
			binData, err := m.c.Get("/lol-game-data/assets/v1/perks.json")
			if err != nil {
				return errs.UnknownError(err)
			}
			array := []model.Perk{}
			err = json.Unmarshal(binData, &array)
			if err != nil {
				return errs.UnknownError(err)
			}
			for i := range array {
				m.Assets.Perks[array[i].Id] = array[i]
			}
			return
		},
		func() (err error) {
			binData, err := m.c.Get("/lol-game-data/assets/v1/items.json")
			if err != nil {
				return errs.UnknownError(err)
			}
			array := []model.Item{}
			err = json.Unmarshal(binData, &array)
			if err != nil {
				return errs.UnknownError(err)
			}
			for i := range array {
				m.Assets.Items[array[i].Id] = array[i]
			}
			m.Assets.Items[0] = model.Item{
				Id: 0,
			}
			return
		},
		func() (err error) {
			binData, err := m.c.Get("/lol-game-data/assets/v1/summoner-spells.json")
			if err != nil {
				return errs.UnknownError(err)
			}
			array := []model.Spell{}
			err = json.Unmarshal(binData, &array)
			if err != nil {
				return errs.UnknownError(err)
			}
			for i := range array {
				m.Assets.Spells[array[i].Id] = array[i]
			}
			return
		},
	)

	for i := 0; i < num; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
