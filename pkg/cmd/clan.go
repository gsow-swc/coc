package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

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
			Name:    c1.Name,
			Tag:     c1.Tag,
			Members: c1.Members,
			Wins:    c1.WarWins,
			Losses:  c1.WarLosses,
			Draws:   c1.WarTies,
			Level:   c1.ClanLevel,
			League:  c1.WarLeague.Name,
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
		Name:    c1.Name,
		Tag:     c1.Tag,
		Members: c1.Members,
		Wins:    c1.WarWins,
		Losses:  c1.WarLosses,
		Draws:   c1.WarTies,
		Level:   c1.ClanLevel,
		League:  c1.WarLeague.Name,
	}
	clans.Clans = append(clans.Clans, c2)
	fmt.Println(clans)
	return nil
}

// ClanMembers gets information about members of a clan
func ClanMembers(c *cli.Context) error {
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

	warClan := warClan{}

	for i, p := range players {
		heroes := getHeroes(p.Heroes)
		m := warClanMember{
			MapPosition:   i + 1,
			Name:          p.Name,
			TownHall:      p.TownHallLevel,
			BarbarianKing: heroes.bk,
			ArcherQueen:   heroes.aq,
			GrandWarden:   heroes.gw,
			RoyalChampion: heroes.rc,
			League:        p.League.Name,
		}

		warClan.Members = append(warClan.Members, m)
	}

	fmt.Println(warClan)

	return nil
}
