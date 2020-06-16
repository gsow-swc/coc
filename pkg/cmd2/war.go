package cmd2

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// warSummaries is the set of summary data for a list of wars
type warSummaries struct {
	wars []warSummary
}

// warSummary is a set of summary data for a clan war
type warSummary struct {
	endTime  time.Time      // The time the war ends, or ended
	teamSize int            // Number of members in the war
	result   string         // Result of the war
	clan     warSummaryClan // A clan in the war
	opponent warSummaryClan // The clan's opponent in the war
}

// warSummaryClan is a set of summary data for a clan in a clan war.
type warSummaryClan struct {
	name                  string  // Name of the clan
	tag                   string  // Tag of the clan
	attacksMade           int     // Number of attacks made in the war
	totalAttacks          int     // Total number of attacks available
	destructionPercentage float32 // Average percentage of the opponents bases destroyed
	stars                 int     // Number of stars for all members in the war
}

// warMap represents the two clans and their members who are in a war.
type warMap struct {
	endTime  time.Time  // The time the war ends, or ended
	teamSize int        // Number of members in the war
	clan     warMapClan // A clan in the war
	opponent warMapClan // The clan's opponent in the war
}

// warMapRoster is the data for one clan in a war.
type warMapClan struct {
	name    string             // Name of the clan
	tag     string             // Tag of the clan
	members []warMapClanMember // Members of the clan in the war
}

// warMapRoster a member of a clan in a war.
type warMapClanMember struct {
	mapPosition   int    // Map position for the member
	name          string // Name of the member
	tag           string // Tag of the member
	townHall      int    // Town hall level
	barbarianKing int    // Barbarian King level
	archerQueen   int    // Archer Queen level
	grandWarden   int    // Grand warden level
	royalChampion int    // Royal champion level
	league        string // League the player is in
}

// warStatus gets the status of a war
type warStatus struct {
	endTime  time.Time     // The time the war ends, or ended
	clan     warStatusClan // Clan in the war
	opponent warStatusClan // The clan's opponent in the war
}

// warStatusClan is the status data for a clan in the war
type warStatusClan struct {
	name    string            // Name of the clan membmer
	tag     string            // Tag of the clan member
	attacks []warStatusAttack // List of attacks by the clan member
	defense warStatusAttack   // Best attack against the clan member
}

// warStatusAttack represents an attack in a clan war
type warStatusAttack struct {
	attacker              string // Name of the attacker
	attackerTownHall      int    // Attacker town hall level
	defender              string // Name of the defender
	defenderTownHall      int    // Town hall level for the defender
	mapPosition           int    // Map position for the attacker
	stars                 int    // Number of stars from the attack
	destructionPercentage int    // Percentage of destruction from the attack
}

// warList is a list of wars a clan has participated in
type warList struct {
	clan    warStatusClan   // The clan
	results []warListResult // The results of each war
}

type warListResult struct {
	endTime                    time.Time     // The time the war ends, or ended
	teamSize                   int           // The number of members in the war
	result                     string        // The result of the war
	clanStars                  int           // The number of stars for the clan
	clanDestructionPercent     float32       // The percentage of destruction for the clan
	opponent                   warStatusClan // The clan's opponent
	opponentStars              int           // The number of stars for the opponent
	opponentDestructionPercent float32       // The opponent destruction percentage
}

// warClan is a clan in a war
type warClan struct {
	endTime      time.Time       // The time the war ends, or ended
	cwlWar       bool            // The war is a CWL war or a regular war
	clanName     string          // The name of the clan
	clanTag      string          // The tag for the clan
	opponentName string          // The name of the clan's opponent
	opponentTag  string          // The tag of the clan's opponent
	members      []warClanMember // The members of the clan in the war
}

// warClanMember is a clan member in the war
type warClanMember struct {
	name        string          // The name of the player
	tag         string          // The tag of the player
	townHall    int             // The town hall level
	mapPosition int             // The map position
	attacks     []warClanAttack // The attacks by the player
}

// warClanAttack is an attack by a member in the war
type warClanAttack struct {
	targetName            string // The name of the target of the attack
	targetTag             string // The tag of the target of the attack
	targetMapPosition     int    // The map position of the target of the attack
	targetTownHall        int    // The town hall level of the target of the attack
	stars                 int    // The number of stars in the attack
	destructionPercentage int    // The destruction percentage of the attack
}

// warTargets are the bases left in a war that haven't been cleared
type warTargets struct {
	endTime      time.Time         // The time the war ends, or ended
	clanName     string            // The name of the clan making the attacks
	clanTag      string            // The tag of the clan making the attacks
	opponentName string            // The name of the clan's opponent
	opponentTag  string            // The tag of the clan's opponent
	targets      []warStatusAttack // The list of non-cleared bases in the war
}

// getWarSummary returns a warSummary created from a clan war
func getWarSummary(w response.ClanWar) warSummary {
	clan := warSummaryClan{
		name:                  w.Clan.Name,
		tag:                   w.Clan.Tag,
		attacksMade:           w.Clan.Attacks,
		totalAttacks:          w.TeamSize * 2,
		destructionPercentage: w.Clan.DestructionPercentage,
		stars:                 w.Clan.Stars,
	}
	opponent := warSummaryClan{
		name:                  w.Opponent.Name,
		tag:                   w.Opponent.Tag,
		attacksMade:           w.Opponent.Attacks,
		totalAttacks:          w.TeamSize * 2,
		destructionPercentage: w.Opponent.DestructionPercentage,
		stars:                 w.Opponent.Stars,
	}
	ws := warSummary{
		endTime:  getTime(w.EndTime),
		result:   w.Result,
		teamSize: w.TeamSize,
		clan:     clan,
		opponent: opponent,
	}
	return ws
}

// WarList gets the list of wars a clan has participated in
func WarList(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the clan wars
	req := request.ClanWars{Tag: tag}
	wars, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	// Create the list of war data
	warList := warSummaries{}
	for _, w := range wars {
		if w.Opponent.Name != "" {
			warList.wars = append(warList.wars, getWarSummary(w))
		}
	}

	fmt.Println(warList)

	return nil
}

// WarCurrent gets the war a clan is participating in
func WarCurrent(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	req := request.ClanCurrentWar{Tag: tag}
	w, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	s := getWarSummary(w)
	fmt.Println(s)

	return nil
}

// WarRoster gets details about the current war for a clan
func WarRoster(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	req := request.ClanCurrentWar{Tag: tag}
	war, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	// Sort the members based on their map position
	sort.Slice(war.Clan.Members, func(i, j int) bool { return war.Clan.Members[i].MapPosition < war.Clan.Members[j].MapPosition })
	sort.Slice(war.Opponent.Members, func(i, j int) bool { return war.Opponent.Members[i].MapPosition < war.Opponent.Members[j].MapPosition })

	// Get the two clans
	var clan response.ClanWarTeam
	var opponent response.ClanWarTeam
	if war.Clan.Tag == tag {
		clan = war.Clan
		opponent = war.Opponent
	} else {
		clan = war.Opponent
		opponent = war.Clan
	}

	wm := warMap{
		endTime:  getTime(war.EndTime),
		teamSize: war.TeamSize,
		clan:     warMapClan{name: clan.Name, tag: clan.Tag},
		opponent: warMapClan{name: opponent.Name, tag: opponent.Tag},
	}

	for i, m := range clan.Members {
		req := request.Player{Tag: m.Tag}
		p, err := req.Get()
		if err != nil {
			log.Error("failed to get the response")
			fmt.Println(err)
			return err
		}

		heroes := getHeroes(p.Heroes)
		cm := warMapClanMember{
			mapPosition:   i + 1,
			name:          m.Name,
			tag:           m.Tag,
			townHall:      m.TownhallLevel,
			barbarianKing: heroes.bk,
			archerQueen:   heroes.aq,
			grandWarden:   heroes.gw,
			royalChampion: heroes.rc,
		}

		wm.clan.members = append(wm.clan.members, cm)
	}

	for i, m := range opponent.Members {
		req := request.Player{Tag: m.Tag}
		p, err := req.Get()
		if err != nil {
			log.Error("failed to get the response")
			fmt.Println(err)
			return err
		}

		heroes := getHeroes(p.Heroes)
		cm := warMapClanMember{
			mapPosition:   i + 1,
			name:          m.Name,
			tag:           m.Tag,
			townHall:      m.TownhallLevel,
			barbarianKing: heroes.bk,
			archerQueen:   heroes.aq,
			grandWarden:   heroes.gw,
			royalChampion: heroes.rc,
		}

		wm.opponent.members = append(wm.opponent.members, cm)
	}

	fmt.Println(wm)

	return nil
}

// String returns a string representation of a list of wars
func (s warSummaries) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"Opponent", "Tag", "Size", "Result", "Stars", "Percent", "OppStars", "OppPercent"})
	for _, w := range s.wars {
		// Get the percentages to two digit precision
		cPercent := float32(int(w.clan.destructionPercentage*100)) / 100
		oPercent := float32(int(w.opponent.destructionPercentage*100)) / 100
		t.AppendRow(table.Row{w.opponent.name, w.opponent.tag, w.teamSize, w.result, w.clan.stars, cPercent, w.opponent.stars, oPercent})
	}

	return t.Render()
}

// String returns a string reprepsentation of a single war
func (s warSummary) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignCenter},
		{Number: 3, Align: text.AlignLeft},
	})

	t.AppendHeader(table.Row{s.clan.name, "", s.opponent.name})
	cStars := strconv.Itoa(s.clan.stars) + "/" + strconv.Itoa(s.teamSize*3)
	oStars := strconv.Itoa(s.opponent.stars) + "/" + strconv.Itoa(s.teamSize*3)
	t.AppendRow(table.Row{cStars, "stars", oStars})
	cPercent := fmt.Sprintf("%.1f", s.clan.destructionPercentage)
	oPercent := fmt.Sprintf("%.1f", s.opponent.destructionPercentage)
	t.AppendRow(table.Row{cPercent, "%", oPercent})
	cAttacks := strconv.Itoa(s.clan.attacksMade) + "/" + strconv.Itoa(s.clan.totalAttacks)
	oAttacks := strconv.Itoa(s.opponent.attacksMade) + "/" + strconv.Itoa(s.opponent.totalAttacks)
	t.AppendRow(table.Row{cAttacks, "attacks", oAttacks})
	t.SetCaption(getTimeLeft(s.endTime))

	return t.Render()
}

// String returns a string representation of two clans that are in a war
func (w warMap) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetTitle(w.clan.name + " vs " + w.opponent.name + " (" + w.opponent.tag + ")")
	t.AppendHeader(table.Row{"#", "Name", "TH", "BK", "AQ", "GW", "RC", "", "Name", "TH", "BK", "AQ", "GW", "RC"})
	for i := range w.clan.members {
		m := w.clan.members[i]
		o := w.opponent.members[i]
		t.AppendRow(table.Row{i + 1, m.name, m.townHall, m.barbarianKing, m.archerQueen, m.grandWarden, m.royalChampion, "   ", o.name, o.townHall, o.barbarianKing, o.archerQueen, o.grandWarden, o.royalChampion})
	}

	return t.Render()
}

// String returns a string representation of a clan that is in a war
func (wc warMapClan) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"#", "Name", "TH", "BK", "AQ", "GW", "RC", "League"})
	for i, m := range wc.members {
		var league string
		if m.league != "" {
			league = m.league
		} else {
			league = "Unranked"
		}
		t.AppendRow(table.Row{i + 1, m.name, m.townHall, m.barbarianKing, m.archerQueen, m.grandWarden, m.royalChampion, league})
	}

	return t.Render()
}

// String returns a string representation of a status of the war
func (ws warStatus) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})

	t.SetTitle(ws.clan.name + " vs " + ws.opponent.name + " (" + ws.opponent.tag + ")")
	t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "Attacker"})
	for _, target := range ws.clan.attacks {
		if target.attacker == "" {
			t.AppendRow(table.Row{target.mapPosition, target.attacker, target.attackerTownHall})
		} else {
			t.AppendRow(table.Row{target.mapPosition, target.attacker, target.attackerTownHall, getStars(target.stars), strconv.Itoa(target.destructionPercentage) + "%", target.attackerTownHall, target.attacker})
		}
	}

	return t.Render()
}

// String provides a string representation of the list of wars a clan has participated in
func (wl warList) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"Opponent", "Opp Tag", "Size", "Result", "Stars", "Percent", "OppStars", "OppPercent"})
	for _, r := range wl.results {
		// Get the percentages to two digit precision
		cPercent := float32(int(r.clanDestructionPercent*100)) / 100
		oPercent := float32(int(r.opponentDestructionPercent*100)) / 100
		t.AppendRow(table.Row{r.opponent.name, r.opponent.tag, r.teamSize, r.result, r.clanStars, cPercent, r.opponentStars, oPercent})
	}

	return t.Render()
}

// String provides a string representation of a clan that is in a war
func (wc warClan) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter},
		{Number: 11, Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})

	t.SetTitle(wc.clanName + " vs " + wc.opponentName + " (" + wc.opponentTag + ")")
	// In a CWL war, you get one attack; in a regular war, you get two attacks
	if wc.cwlWar {
		t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "#", "Attacked"})
	} else {
		t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "#", "Attacked", "", threestar, "%", "TH", "#", "Attacked"})
	}

	for _, m := range wc.members {
		if len(m.attacks) == 0 {
			t.AppendRow(table.Row{m.mapPosition, m.name, m.townHall})
		} else if len(m.attacks) == 1 {
			a := m.attacks[0]
			t.AppendRow(table.Row{m.mapPosition, m.name, m.townHall, getStars(a.stars), strconv.Itoa(a.destructionPercentage) + "%", a.targetTownHall, a.targetMapPosition, a.targetName})
		} else {
			a1 := m.attacks[0]
			a2 := m.attacks[1]
			t.AppendRow(table.Row{m.mapPosition, m.name, m.townHall, a1.targetName, a1.targetMapPosition, a1.targetTownHall, getStars(a1.stars), strconv.Itoa(a1.destructionPercentage) + "%", "   ", a2.targetName, a2.targetMapPosition, a2.targetTownHall, getStars(a2.stars), strconv.Itoa(a2.destructionPercentage) + "%"})
		}
	}

	t.SetCaption(getTimeLeft(wc.endTime))

	return t.Render()
}

// String returns a string represntation of the non-cleared targets in a war
func (wt warTargets) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})

	t.SetTitle(wt.clanName + " vs " + wt.opponentName + " (" + wt.opponentTag + ")")
	t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%"})
	for _, target := range wt.targets {
		if target.destructionPercentage == 0 {
			t.AppendRow(table.Row{target.mapPosition, target.defender, target.defenderTownHall})
		} else {
			t.AppendRow(table.Row{target.mapPosition, target.defender, target.defenderTownHall, getStars(target.stars), strconv.Itoa(target.destructionPercentage) + "%"})
		}
	}

	return t.Render()
}
