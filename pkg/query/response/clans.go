package response

import "encoding/json"

// Clan is a clan in Clash of Clans.
type Clan struct {
	WarLeague struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"warLeague"`
	MemberList       []ClanMember `json:"memberList"`
	WarLosses        int          `json:"warLosses"`
	Tag              string       `json:"tag"`
	ClanVersusPoints int          `json:"clanVersusPoints"`
	RequiredTrophies int          `json:"requiredTrophies"`
	WarFrequency     string       `json:"warFrequency"`
	ClanLevel        int          `json:"clanLevel"`
	WarWinStreak     int          `json:"warWinStreak"`
	WarWins          int          `json:"warWins"`
	ClanPoints       int          `json:"clanPoints"`
	IsWarLogPublic   bool         `json:"isWarLogPublic"`
	WarTies          int          `json:"warTies"`
	Labels           []struct {
		Name     string   `json:"name"`
		ID       int      `json:"id"`
		IconUrls IconUrls `json:"iconUrls"`
	} `json:"labels"`
	Name     string `json:"name"`
	Location struct {
		LocalizedName string `json:"localizedName"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		IsCountry     bool   `json:"isCountry"`
		CountryCode   string `json:"countryCode"`
	} `json:"location"`
	Type        string    `json:"type"`
	Members     int       `json:"members"`
	Description string    `json:"description"`
	BadgeUrls   BadgeUrls `json:"badgeUrls"`
}

// String returns a string representation of a clan
func (c Clan) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanMember is a member of a given clan.
type ClanMember struct {
	League struct {
		Name     string   `json:"name"`
		ID       int      `json:"id"`
		IconUrls IconUrls `json:"iconUrls"`
	} `json:"league"`
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              string `json:"role"`
	ExpLevel          int    `json:"expLevel"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
}

// String returns a string representation of a clan member
func (m ClanMember) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// ClanWar is a given war in a clan's war log.
type ClanWar struct {
	State                string      `json:"state,omitempty"`
	TeamSize             int         `json:"teamSize"`
	PreparationStartTime string      `json:"preparationStartTime,omitempty"`
	StartTime            string      `json:"startTime,omitempty"`
	EndTime              string      `json:"endTime,omitempty"`
	Result               string      `json:"result,omitempty"`
	Clan                 ClanWarTeam `json:"clan"`
	Opponent             ClanWarTeam `json:"opponent"`
}

// String returns a string representation of a clan war
func (cw ClanWar) String() string {
	b, _ := json.Marshal(cw)
	return string(b)
}

// ClanWarTeam is the clan that is participating in the clan war.
type ClanWarTeam struct {
	Attacks               int             `json:"attacks"`
	BadgeUrls             BadgeUrls       `json:"badgeUrls"`
	ClanLevel             int             `json:"clanLevel"`
	DestructionPercentage float32         `json:"destructionPercentage"`
	ExpEarned             int             `json:"expEarned"`
	Members               []ClanWarMember `json:"members,omitempty"`
	Name                  string          `json:"name"`
	Stars                 int             `json:"stars"`
	Tag                   string          `json:"tag"`
}

// String returns a string representation of a clan war team
func (cwt ClanWarTeam) String() string {
	b, _ := json.Marshal(cwt)
	return string(b)
}

// ClanWarMember is a member who participated in a clan war.
type ClanWarMember struct {
	Attacks            []ClanWarAttack `json:"attacks,omitempty"`
	BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack,omitempty"`
	MapPosition        int             `json:"mapPosition"`
	Name               string          `json:"name"`
	OpponentAttacks    int             `json:"opponentAttacks"`
	Tag                string          `json:"tag"`
	TownhallLevel      int             `json:"townhallLevel"`
}

// String returns a string representation of a clan war member
func (cwm ClanWarMember) String() string {
	b, _ := json.Marshal(cwm)
	return string(b)
}

// ClanWarAttack is an attack made in a clan war.
type ClanWarAttack struct {
	Order                 int    `json:"order"`
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
}

// String returns a string representation of a clan war atack
func (cwa ClanWarAttack) String() string {
	b, _ := json.Marshal(cwa)
	return string(b)
}

// ClanWarLeagueGroup is a clan's current clan war league group.
type ClanWarLeagueGroup struct {
	Tag    string `json:"tag"`
	State  string `json:"state"`
	Season string `json:"season"`
	Clans  []struct {
		Tag       string `json:"tag"`
		ClanLevel int    `json:"clanLevel"`
		Name      string `json:"name"`
		Members   []struct {
			Tag           string `json:"tag"`
			TownHallLevel int    `json:"townHallLevel"`
			Name          string `json:"name"`
		} `json:"members"`
		BadgeUrls BadgeUrls `json:"badgeUrls"`
	} `json:"clans"`
	Rounds []struct {
		WarTags []string `json:"warTags"`
	} `json:"rounds"`
}

// String returns a string representation of a clan war league group
func (lg ClanWarLeagueGroup) String() string {
	b, _ := json.Marshal(lg)
	return string(b)
}

// ClanWarLeagueWar is information about an individual clan war league war
type ClanWarLeagueWar struct {
	Clan                 ClanWarTeam `json:"clan"`
	EndTime              string      `json:"endTime"`
	Opponent             ClanWarTeam `json:"opponent"`
	PreparationStartTime string      `json:"preparationStartTime"`
	StartTime            string      `json:"startTime"`
	State                string      `json:"state"`
	TeamSize             int         `json:"teamSize"`
	WarStartTime         string      `json:"warStartTime"`
}

// String returns a string representation of a clan war league war
func (lw ClanWarLeagueWar) String() string {
	b, _ := json.Marshal(lw)
	return string(b)
}
