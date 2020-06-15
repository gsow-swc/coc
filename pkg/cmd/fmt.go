package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	star      = "⭐"
	zerostar  = ""
	onestar   = star
	twostar   = star + star
	threestar = star + star + star
)

func getStars(stars int) string {
	if stars == 0 {
		return zerostar
	} else if stars == 1 {
		return onestar
	} else if stars == 2 {
		return twostar
	}

	return threestar
}

type clans struct {
	Clans []clan
}

func (cs clans) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"Name", "Tag", "Members", "Wins", "Losses", "Draws", "Level", "League"})
	for _, c := range cs.Clans {
		t.AppendRow(table.Row{c.Name, c.Tag, c.Members, c.Wins, c.Losses, c.Draws, c.Level, c.League})
	}
	return t.Render()
}

type clan struct {
	Name    string
	Tag     string
	Members int
	Wins    int
	Losses  int
	Draws   int
	Level   int
	League  string
}

func (c clan) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"Name", "Tag", "Members", "Wins", "Losses", "Draws", "Level", "League"})
	t.AppendRow(table.Row{c.Name, c.Tag, c.Members, c.Wins, c.Losses, c.Draws, c.Level, c.League})

	return t.Render()
}

// Scoreboard holds a war scoreboard for a clan and their opponent
type scoreboard struct {
	TeamSize int
	EndTime  time.Time
	Clan     scoreboardClan
	Opponent scoreboardClan
}

// ScoreboardClan holds the scoreboard for one clan
type scoreboardClan struct {
	Name                  string
	AttacksMade           int
	TotalAttacks          int
	DestructionPercentage float32
	Stars                 int
}

// String returns a string reprepsentation of a scoreboard
func (sb scoreboard) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignCenter},
		{Number: 3, Align: text.AlignLeft},
	})

	t.AppendHeader(table.Row{sb.Clan.Name, "", sb.Opponent.Name})
	cStars := strconv.Itoa(sb.Clan.Stars) + "/" + strconv.Itoa(sb.TeamSize*3)
	oStars := strconv.Itoa(sb.Opponent.Stars) + "/" + strconv.Itoa(sb.TeamSize*3)
	t.AppendRow(table.Row{cStars, "stars", oStars})
	cPercent := fmt.Sprintf("%.1f", sb.Clan.DestructionPercentage)
	oPercent := fmt.Sprintf("%.1f", sb.Opponent.DestructionPercentage)
	t.AppendRow(table.Row{cPercent, "%", oPercent})
	cAttacks := strconv.Itoa(sb.Clan.AttacksMade) + "/" + strconv.Itoa(sb.Clan.TotalAttacks)
	oAttacks := strconv.Itoa(sb.Opponent.AttacksMade) + "/" + strconv.Itoa(sb.Opponent.TotalAttacks)
	t.AppendRow(table.Row{cAttacks, "attacks", oAttacks})
	t.SetCaption(getTimeLeft(sb.EndTime))

	return t.Render()
}

// WarMap holds information about two clans that are in war or preparing for a war
type warMap struct {
	Clan     warClan
	Opponent warClan
}

// WarClan is one clan that is in or preparing for a war
type warClan struct {
	Name    string
	Tag     string
	Members []warClanMember
}

// WarClanMember is a member of a clan in or preparing for a war
type warClanMember struct {
	MapPosition   int
	Name          string
	TownHall      int
	BarbarianKing int
	ArcherQueen   int
	GrandWarden   int
	RoyalChampion int
	League        string
}

// String returns a string representation of a clan that is in a war
func (wc warClan) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"#", "Name", "TH", "BK", "AQ", "GW", "RC", "League"})
	for i, m := range wc.Members {
		var league string
		if m.League != "" {
			league = m.League
		} else {
			league = "Unranked"
		}
		t.AppendRow(table.Row{i + 1, m.Name, m.TownHall, m.BarbarianKing, m.ArcherQueen, m.GrandWarden, m.RoyalChampion, league})
	}

	return t.Render()
}

// String returns a string representation of two clans that are in a war
func (w warMap) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetTitle(w.Clan.Name + " vs " + w.Opponent.Name + " (" + w.Opponent.Tag + ")")
	t.AppendHeader(table.Row{"#", "Name", "TH", "BK", "AQ", "GW", "RC", "", "Name", "TH", "BK", "AQ", "GW", "RC"})
	for i := range w.Clan.Members {
		m := w.Clan.Members[i]
		o := w.Opponent.Members[i]
		t.AppendRow(table.Row{i + 1, m.Name, m.TownHall, m.BarbarianKing, m.ArcherQueen, m.GrandWarden, m.RoyalChampion, "   ", o.Name, o.TownHall, o.BarbarianKing, o.ArcherQueen, o.GrandWarden, o.RoyalChampion})
	}

	return t.Render()
}

// WarStatus stores information about an ongoing war
type warStatus struct {
	ClanName        string
	OpponentName    string
	Targets         []warStatusTarget
	OpponentTargets []warStatusTarget
}

// WarStatusTarget stores information about the member of the other clan in a war
type warStatusTarget struct {
	Name                  string
	Attacker              string
	AttackerTownHall      int
	MapPosition           int
	Stars                 int
	TownHall              int
	DestructionPercentage int
}

// String returns a string represntation of a summary of the war
func (ws warStatus) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})

	t.SetTitle(ws.ClanName + " vs " + ws.OpponentName)
	t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "Attacker"})
	for _, target := range ws.Targets {
		if target.Attacker == "" {
			t.AppendRow(table.Row{target.MapPosition, target.Name, target.TownHall})
		} else {
			t.AppendRow(table.Row{target.MapPosition, target.Name, target.TownHall, getStars(target.Stars), strconv.Itoa(target.DestructionPercentage) + "%", target.AttackerTownHall, target.Attacker})
		}
	}

	return t.Render()
}

// wars is a list of wars a clan has participated in
type wars struct {
	Results []warResult
}

// warResult is a summary list of the results of one or more wars
type warResult struct {
	OpponentName    string
	OpponentTag     string
	TeamSize        int
	Result          string
	Stars           int
	Percent         float32
	OpponentStars   int
	OpponentPercent float32
}

// String provides a string representation of the list of wars a clan has participated in
func (w wars) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"Opponent", "Opp Tag", "Size", "Result", "Stars", "Percent", "OppStars", "OppPercent"})
	for _, r := range w.Results {
		// Get the percentages to two digit precision
		cPercent := float32(int(r.Percent*100)) / 100
		oPercent := float32(int(r.OpponentPercent*100)) / 100
		t.AppendRow(table.Row{r.OpponentName, r.OpponentTag, r.TeamSize, r.Result, r.Stars, cPercent, r.OpponentStars, oPercent})
	}

	return t.Render()
}

type war2 struct {
	Clan     war2Clan
	Opponent war2Clan
	CwlWar   bool
	EndTime  time.Time
}

type war2Clan struct {
	Name    string
	Tag     string
	Members []war2ClanMember
}

type war2ClanMember struct {
	Name        string
	Tag         string
	TownHall    int
	MapPosition int
	Attacks     []war2ClanAttack
}

type war2ClanAttack struct {
	TargetName            string
	TargetTag             string
	TargetMapPosition     int
	TargetTownhall        int
	Stars                 int
	DestructionPercentage int
}

func (w war2) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter},
		{Number: 11, Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})

	t.SetTitle(w.Clan.Name + " vs " + w.Opponent.Name)
	// In a CWL war, you get one attack; in a regular war, you get two attacks
	if w.CwlWar {
		t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "#", "Attacked"})
	} else {
		t.AppendHeader(table.Row{"#", "Name", "TH", threestar, "%", "TH", "#", "Attacked", "", threestar, "%", "TH", "#", "Attacked"})
	}

	for _, m := range w.Clan.Members {
		if len(m.Attacks) == 0 {
			t.AppendRow(table.Row{m.MapPosition, m.Name, m.TownHall})
		} else if len(m.Attacks) == 1 {
			a := m.Attacks[0]
			t.AppendRow(table.Row{m.MapPosition, m.Name, m.TownHall, getStars(a.Stars), strconv.Itoa(a.DestructionPercentage) + "%", a.TargetTownhall, a.TargetMapPosition, a.TargetName})
		} else {
			a1 := m.Attacks[0]
			a2 := m.Attacks[0]
			t.AppendRow(table.Row{m.MapPosition, m.Name, m.TownHall, a1.TargetName, a1.TargetMapPosition, a1.TargetTownhall, getStars(a1.Stars), strconv.Itoa(a1.DestructionPercentage) + "%", "   ", a2.TargetName, a2.TargetMapPosition, a2.TargetTownhall, getStars(a2.Stars), strconv.Itoa(a2.DestructionPercentage) + "%"})
		}
	}

	t.SetCaption(getTimeLeft(w.EndTime))

	return t.Render()
}
