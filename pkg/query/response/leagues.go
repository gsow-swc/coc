package response

import "encoding/json"

// League lists leagues
type League struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	IconUrls struct {
	} `json:"iconUrls"`
}

// String returns a string representation of a league
func (l League) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LeagueSeason is a league season.
type LeagueSeason struct {
	ID string `json:"id"`
}

// String returns a string representation of a league season
func (ls LeagueSeason) String() string {
	b, _ := json.Marshal(ls)
	return string(b)
}

// LeagueSeasonRanking is the league season ranking.
type LeagueSeasonRanking struct {
	Clan struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		BadgeUrls struct {
		} `json:"badgeUrls"`
	} `json:"clan"`
	League       League `json:"league"`
	AttackWins   int    `json:"attackWins"`
	DefenseWins  int    `json:"defenseWins"`
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	ExpLevel     int    `json:"expLevel"`
	Rank         int    `json:"rank"`
	PreviousRank int    `json:"previousRank"`
	Trophies     int    `json:"trophies"`
}

// String returns a string representation of a league season
func (lsr LeagueSeasonRanking) String() string {
	b, _ := json.Marshal(lsr)
	return string(b)
}

// WarLeague is information about a war league.
type WarLeague struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// String returns a string representation of a war league
func (wl WarLeague) String() string {
	b, _ := json.Marshal(wl)
	return string(b)
}
