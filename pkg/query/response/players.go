package response

import "encoding/json"

// Player is a single player in Clash of Clans.
type Player struct {
	Tag                string `json:"tag"`
	Name               string `json:"name"`
	TownHallLevel      int    `json:"townHallLevel"`
	ExpLevel           int    `json:"expLevel"`
	Trophies           int    `json:"trophies"`
	BestTrophies       int    `json:"bestTrophies"`
	WarStars           int    `json:"warStars"`
	AttackWins         int    `json:"attackWins"`
	DefenseWins        int    `json:"defenseWins"`
	BuilderHallLevel   int    `json:"builderHallLevel"`
	VersusTrophies     int    `json:"versusTrophies"`
	BestVersusTrophies int    `json:"bestVersusTrophies"`
	VersusBattleWins   int    `json:"versusBattleWins"`
	Role               string `json:"role"`
	Donations          int    `json:"donations"`
	DonationsReceived  int    `json:"donationsReceived"`
	Clan               struct {
		Tag       string    `json:"tag"`
		Name      string    `json:"name"`
		ClanLevel int       `json:"clanLevel"`
		BadgeUrls BadgeUrls `json:"badgeUrls"`
	} `json:"clan"`
	League struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		IconUrls IconUrls `json:"iconUrls"`
	} `json:"league"`
	Achievements []struct {
		Name           string `json:"name"`
		Stars          int    `json:"stars"`
		Value          int    `json:"value"`
		Target         int    `json:"target"`
		Info           string `json:"info"`
		CompletionInfo string `json:"completionInfo"`
		Village        string `json:"village"`
	} `json:"achievements"`
	VersusBattleWinCount int `json:"versusBattleWinCount"`
	Labels               []struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		IconUrls IconUrls `json:"iconUrls"`
	} `json:"labels"`
	Troops []Troop `json:"troops"`
	Heroes []Troop `json:"heroes"`
	Spells []Troop `json:"spells"`
}

// Troop represents a troop, hero or spell in Clash of Clans
type Troop struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
}

// String returns a string representation of a player
func (p Player) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
