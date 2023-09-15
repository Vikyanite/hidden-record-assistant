package model

type Assets struct {
	Champions map[int]Champion  `json:"champions"`
	Queues    map[int]Queue     `json:"queues"`
	Perks     map[int]Perk      `json:"perks"`
	Items     map[int]Item      `json:"items"`
	Spells    map[int]Spell     `json:"spells"`
	PerkStyle map[int]PerkStyle `json:"perkStyles"`
	Tier      map[string]Tier   `json:"tier"`
}

func (a *Assets) GetChampionById(id int) Champion {
	return a.Champions[id]
}

func (a *Assets) GetQueueById(id int) Queue {
	return a.Queues[id]
}

func (a *Assets) GetPerkById(id int) Perk {
	return a.Perks[id]
}

func (a *Assets) GetItemById(id int) Item {
	return a.Items[id]
}

func (a *Assets) GetSpellById(id int) Spell {
	return a.Spells[id]
}

func (a *Assets) GetPerkStyleById(id int) PerkStyle {
	return a.PerkStyle[id]
}

type PerkStyle struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IconPath    string `json:"iconPath"`
}

type Perk struct {
	Id                       int      `json:"id"`
	Name                     string   `json:"name"`
	MajorChangePatchVersion  string   `json:"majorChangePatchVersion"`
	Tooltip                  string   `json:"tooltip"`
	ShortDesc                string   `json:"shortDesc"`
	LongDesc                 string   `json:"longDesc"`
	RecommendationDescriptor string   `json:"recommendationDescriptor"`
	IconPath                 string   `json:"iconPath"`
	EndOfGameStatDescs       []string `json:"endOfGameStatDescs"`
	RecommendationAttributes struct{} `json:"recommendationDescriptorAttributes"`
}

type Spell struct {
	Id            int      `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	SummonerLevel int      `json:"summonerLevel"`
	Cooldown      int      `json:"cooldown"`
	GameModes     []string `json:"gameModes"`
	IconPath      string   `json:"iconPath"`
}

type Queue struct {
	Name                string `json:"name"`
	ShortName           string `json:"shortName"`
	Description         string `json:"description"`
	Id                  int    `json:"id"`
	DetailedDescription string `json:"detailedDescription"`
}

type Champion struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Alias              string   `json:"alias"`
	SquarePortraitPath string   `json:"squarePortraitPath"`
	Roles              []string `json:"roles"`
}

type Item struct {
	Id                       int      `json:"id"`
	Name                     string   `json:"name"`
	Description              string   `json:"description"`
	Active                   bool     `json:"active"`
	InStore                  bool     `json:"inStore"`
	From                     []int    `json:"from"`
	To                       []int    `json:"to"`
	Categories               []string `json:"categories"`
	MaxStacks                int      `json:"maxStacks"`
	RequiredChampion         string   `json:"requiredChampion"`
	RequiredAlly             string   `json:"requiredAlly"`
	RequiredBuffCurrencyName string   `json:"requiredBuffCurrencyName"`
	RequiredBuffCurrencyCost int      `json:"requiredBuffCurrencyCost"`
	SpecialRecipe            int      `json:"specialRecipe"`
	IsEnchantment            bool     `json:"isEnchantment"`
	Price                    int      `json:"price"`
	PriceTotal               int      `json:"priceTotal"`
	IconPath                 string   `json:"iconPath"`
}

type Tier struct {
	IconPath string `json:"iconPath"`
	Name     string `json:"name"`
	Division string `json:"division"`
}
