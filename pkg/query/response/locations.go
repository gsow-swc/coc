package response

import "encoding/json"

// Location is information about a location
type Location struct {
	LocalizedName string `json:"localizedName"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}

// String returns a string representation of a location
func (l Location) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationClanRanking is the clan ranking for a specific location.
type LocationClanRanking struct {
	ClanLevel  int `json:"clanLevel"`
	ClanPoints int `json:"clanPoints"`
	Location   struct {
		LocalizedName string `json:"localizedName"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		IsCountry     bool   `json:"isCountry"`
		CountryCode   string `json:"countryCode"`
	} `json:"location"`
	Members      int    `json:"members"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	Rank         int    `json:"rank"`
	PreviousRank int    `json:"previousRank"`
	BadgeUrls    struct {
	} `json:"badgeUrls"`
}

// String returns a string representation of a location clan ranking
func (l LocationClanRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationPlayerRanking is the ranking of a player for specific location.
type LocationPlayerRanking struct {
	Clan struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		BadgeUrls struct {
		} `json:"badgeUrls"`
	} `json:"clan"`
	League struct {
		Name     string `json:"name"`
		ID       int    `json:"id"`
		IconUrls struct {
		} `json:"iconUrls"`
	} `json:"league"`
	AttackWins   int    `json:"attackWins"`
	DefenseWins  int    `json:"defenseWins"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	ExpLevel     int    `json:"expLevel"`
	Rank         int    `json:"rank"`
	PreviousRank int    `json:"previousRank"`
	Trophies     int    `json:"trophies"`
}

// String returns a string representation of a location player ranking
func (l LocationPlayerRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationClanVersusRanking is the clan versus ranking for a specific location
type LocationClanVersusRanking struct {
	ClanVersusPoints int `json:"clanVersusPoints"`
	ClanPoints       int `json:"clanPoints"`
}

// String returns a string representation of a clan-versus ranking for a location
func (l LocationClanVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationPlayerVersusRanking is the player ranking for a specific location
type LocationPlayerVersusRanking struct {
	Clan struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		BadgeUrls struct {
		} `json:"badgeUrls"`
	} `json:"clan"`
	VersusBattleWins int    `json:"versusBattleWins"`
	Tag              string `json:"tag"`
	Name             string `json:"name"`
	ExpLevel         int    `json:"expLevel"`
	Rank             int    `json:"rank"`
	PreviousRank     int    `json:"previousRank"`
	VersusTrophies   int    `json:"versusTrophies"`
}

// String returns a string representation of a player-versus ranking for a location
func (l LocationPlayerVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}
