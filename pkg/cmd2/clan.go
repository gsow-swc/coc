package cmd2

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	"github.com/jedib0t/go-pretty/v6/table"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// clans is an overview of a list of clans
type clans struct {
	Clans []clan // List of clans
}

// clan is an overview of a given clan
type clan struct {
	name    string // Name of the clan
	tag     string // Tag for the clan
	members int    // Number of members in the clan
	wins    int    // Number of regular war wins for the clan
	losses  int    // Number of regular war losses for the clan
	draws   int    // Number of regular war draws for the clan
	level   int    // Clan level
	league  string // CWL league for the clan
}

// ClanList lists all clans that match the provided filters
func ClanList(c *cli.Context) error {
	req := request.Clans{
		Name:          c.String("name"),
		WarFrequency:  c.String("frequency"),
		LocationID:    c.Int("location"),
		MinMembers:    c.Int("minmembers"),
		MaxMembers:    c.Int("maxmembers"),
		LabelIDs:      c.StringSlice("labels"),
		MinClanPoints: c.Int("minpoints"),
		MinClanLevel:  c.Int("minlevel"),
	}

	// Make sure at least one filter is specified
	if req.Name == "" && req.WarFrequency == "" && req.LocationID == 0 && req.MinMembers == 0 && req.MaxMembers == 0 && (req.LabelIDs == nil || len(req.LabelIDs) == 0) && req.MinClanPoints == 0 && req.MaxMembers == 0 {
		cli.ShowCommandHelpAndExit(c, "ls", -1)
	}

	cs, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		return err
	}

	clans := clans{}
	for _, c1 := range cs {
		c2 := clan{
			name:    c1.Name,
			tag:     c1.Tag,
			members: c1.Members,
			wins:    c1.WarWins,
			losses:  c1.WarLosses,
			draws:   c1.WarTies,
			level:   c1.ClanLevel,
			league:  c1.WarLeague.Name,
		}
		clans.Clans = append(clans.Clans, c2)
	}
	fmt.Println(clans)

	return nil
}

// ClanGet gets information about a specific clan
func ClanGet(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	req := request.Clan{Tag: tag}
	c1, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	clans := clans{}
	c2 := clan{
		name:    c1.Name,
		tag:     c1.Tag,
		members: c1.Members,
		wins:    c1.WarWins,
		losses:  c1.WarLosses,
		draws:   c1.WarTies,
		level:   c1.ClanLevel,
		league:  c1.WarLeague.Name,
	}
	clans.Clans = append(clans.Clans, c2)
	fmt.Println(clans)
	return nil
}

// ClanMembersGet gets information about members of a clan
func ClanMembersGet(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	req := request.ClanMembers{Tag: tag}
	members, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	var players []response.Player
	for _, member := range members {
		req := request.Player{Tag: member.Tag}
		player, err := req.Get()
		if err != nil {
			log.Error("failed to get the response")
			fmt.Println(err)
			return err
		}
		players = append(players, player)
	}
	// Sort by TH level and by heros and then by name
	sort.Slice(players, func(i, j int) bool {
		var res bool
		if players[i].TownHallLevel != players[j].TownHallLevel {
			res = players[i].TownHallLevel > players[j].TownHallLevel
		} else {
			heroes1 := getHeroes(players[i].Heroes)
			heroes2 := getHeroes(players[j].Heroes)

			if heroes1.rc != heroes2.rc {
				res = heroes1.rc > heroes2.rc
			} else if heroes1.gw != heroes2.gw {
				res = heroes1.gw > heroes2.gw
			} else if heroes1.aq != heroes2.aq {
				res = heroes1.aq > heroes2.aq
			} else if heroes1.bk != heroes2.bk {
				res = heroes1.bk > heroes2.bk
			} else {
				res = strings.ToLower(players[i].Name) < strings.ToLower(players[j].Name)
			}
		}
		return res
	})

	clan := warMapClan{}

	for i, p := range players {
		heroes := getHeroes(p.Heroes)
		m := warMapClanMember{
			mapPosition:   i + 1,
			name:          p.Name,
			townHall:      p.TownHallLevel,
			barbarianKing: heroes.bk,
			archerQueen:   heroes.aq,
			grandWarden:   heroes.gw,
			royalChampion: heroes.rc,
			league:        p.League.Name,
		}

		clan.members = append(clan.members, m)
	}

	fmt.Println(clan)

	return nil
}

// String returns a string represenntation of a list of clans.
func (cs clans) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"Name", "Tag", "Members", "Wins", "Losses", "Draws", "Level", "League"})
	for _, c := range cs.Clans {
		t.AppendRow(table.Row{c.name, c.tag, c.members, c.wins, c.losses, c.draws, c.level, c.league})
	}
	return t.Render()
}

// String returns a string representation of a clan.
func (c clan) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"Name", "Tag", "Members", "Wins", "Losses", "Draws", "Level", "League"})
	t.AppendRow(table.Row{c.name, c.tag, c.members, c.wins, c.losses, c.draws, c.level, c.league})

	return t.Render()
}
