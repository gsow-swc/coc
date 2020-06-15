package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func getWar(tag string, r int) (response.ClanWarLeagueWar, error) {
	var war response.ClanWarLeagueWar

	// Get the Clan War League information for the clan
	req := request.ClanWarLeagueGroup{Tag: tag}
	league, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return war, err
	}

	// Find the current war if a round isn't specified
	if r == 0 {
		for i, round := range league.Rounds {
			r = i + 1
			// War has not started
			if round.WarTags[0] == "#0" {
				r = i
				break
			}
		}
	}

	// Get the array index, ensuring it is 0..6 inclusive
	i := r - 1
	if i < 0 {
		i = 0
	} else if i > 6 {
		i = 6
	}

	// Get the war tags
	var warTags []string
	for _, warTag := range league.Rounds[i].WarTags {
		warTags = append(warTags, warTag)
	}

	// Find the war for the clan we are searching for
	for _, wt := range warTags {
		req := request.ClanWarLeagueWar{Tag: wt}
		war, err = req.Get()
		if err != nil {
			log.Error("failed to get the response")
			fmt.Println(err)
			return war, err
		}

		// We found the clan we are searching for, so exit the loop
		if war.Clan.Tag == tag || war.Opponent.Tag == tag {
			break
		}
	}

	// Sort the members based on their map position
	sort.Slice(war.Clan.Members, func(i, j int) bool { return war.Clan.Members[i].MapPosition < war.Clan.Members[j].MapPosition })
	sort.Slice(war.Opponent.Members, func(i, j int) bool { return war.Opponent.Members[i].MapPosition < war.Opponent.Members[j].MapPosition })

	return war, nil
}

// CwlScoreboard gets summary data about the current CWL war
func CwlScoreboard(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the war
	round := c.Int("round")
	war, err := getWar(tag, round)
	if err != nil {
		return err
	}

	// Create the scoreboard
	var sb scoreboard
	clan1 := scoreboardClan{
		Name:                  war.Clan.Name,
		AttacksMade:           war.Clan.Attacks,
		TotalAttacks:          war.TeamSize,
		DestructionPercentage: war.Clan.DestructionPercentage,
		Stars:                 war.Clan.Stars,
	}
	clan2 := scoreboardClan{
		Name:                  war.Opponent.Name,
		AttacksMade:           war.Opponent.Attacks,
		TotalAttacks:          war.TeamSize,
		DestructionPercentage: war.Opponent.DestructionPercentage,
		Stars:                 war.Opponent.Stars,
	}
	layout := "20060102T150405.000Z"
	t, err := time.Parse(layout, war.EndTime)
	if err != nil {
		return err
	}
	if war.Clan.Tag == tag {
		sb = scoreboard{
			TeamSize: war.TeamSize,
			EndTime:  t,
			Clan:     clan1,
			Opponent: clan2,
		}
	} else {
		sb = scoreboard{
			TeamSize: war.TeamSize,
			EndTime:  t,
			Clan:     clan2,
			Opponent: clan1,
		}
	}

	fmt.Println(sb)

	return nil
}

// CwlAttack gets information about a given CWL war
func CwlAttack(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the war
	round := c.Int("round")
	war, err := getWar(tag, round)
	if err != nil {
		return err
	}

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

	// Build a map of the player tag to the player name
	type player struct {
		name   string
		th     int
		mapPos int
	}
	playerMap := make(map[string]player)
	for i, m := range clan.Members {
		playerMap[m.Tag] = player{name: m.Name, th: m.TownhallLevel, mapPos: i + 1}
	}
	for i, m := range opponent.Members {
		playerMap[m.Tag] = player{name: m.Name, th: m.TownhallLevel, mapPos: i + 1}
	}

	layout := "20060102T150405.000Z"
	warEndTime, err := time.Parse(layout, war.EndTime)
	w := war2{
		Clan: war2Clan{
			Name: clan.Name,
			Tag:  clan.Tag,
		},
		Opponent: war2Clan{
			Name: opponent.Name,
			Tag:  opponent.Tag,
		},
		CwlWar:  true,
		EndTime: warEndTime,
	}
	for i, member := range clan.Members {
		m := war2ClanMember{
			Name:        member.Name,
			Tag:         member.Tag,
			TownHall:    member.TownhallLevel,
			MapPosition: i + 1,
		}
		for _, attack := range member.Attacks {
			p := playerMap[attack.DefenderTag]
			a := war2ClanAttack{
				TargetName:            p.name,
				TargetTag:             attack.DefenderTag,
				TargetMapPosition:     p.mapPos,
				TargetTownhall:        p.th,
				Stars:                 attack.Stars,
				DestructionPercentage: attack.DestructionPercentage,
			}
			m.Attacks = append(m.Attacks, a)
		}
		w.Clan.Members = append(w.Clan.Members, m)
	}

	fmt.Println(w)

	return nil
}

// CwlDefend gets information about the current CWL war
func CwlDefend(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the war
	round := c.Int("round")
	war, err := getWar(tag, round)
	if err != nil {
		return err
	}

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

	// Build a map of the player tag to the player name
	type player struct {
		name string
		th   int
	}
	playerMap := make(map[string]player)
	for _, m := range clan.Members {
		playerMap[m.Tag] = player{name: m.Name, th: m.TownhallLevel}
	}
	for _, m := range opponent.Members {
		playerMap[m.Tag] = player{name: m.Name, th: m.TownhallLevel}
	}

	var ws warStatus
	ws = warStatus{ClanName: clan.Name, OpponentName: opponent.Name}
	for i, m := range clan.Members {
		target := warStatusTarget{
			Name:        m.Name,
			MapPosition: i + 1,
			TownHall:    m.TownhallLevel,
		}
		if m.BestOpponentAttack.AttackerTag != "" {
			attacker := playerMap[m.BestOpponentAttack.AttackerTag]
			target.Attacker = attacker.name
			target.AttackerTownHall = attacker.th
			target.Stars = m.BestOpponentAttack.Stars
			target.DestructionPercentage = m.BestOpponentAttack.DestructionPercentage
		}
		ws.Targets = append(ws.Targets, target)
	}

	fmt.Println(ws)
	return nil
}

// CwlRoster gets information about the next clan to be faced in CWL
func CwlRoster(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the war
	round := c.Int("round")
	war, err := getWar(tag, round)
	if err != nil {
		return err
	}

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
		Clan:     warClan{Name: clan.Name, Tag: clan.Tag},
		Opponent: warClan{Name: opponent.Name, Tag: opponent.Tag},
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
		cm := warClanMember{
			MapPosition:   i + 1,
			Name:          m.Name,
			TownHall:      m.TownhallLevel,
			BarbarianKing: heroes.bk,
			ArcherQueen:   heroes.aq,
			GrandWarden:   heroes.gw,
			RoyalChampion: heroes.rc,
		}

		wm.Clan.Members = append(wm.Clan.Members, cm)
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
		cm := warClanMember{
			MapPosition:   i + 1,
			Name:          m.Name,
			TownHall:      m.TownhallLevel,
			BarbarianKing: heroes.bk,
			ArcherQueen:   heroes.aq,
			GrandWarden:   heroes.gw,
			RoyalChampion: heroes.rc,
		}

		wm.Opponent.Members = append(wm.Opponent.Members, cm)
	}

	fmt.Println(wm)

	return nil
}
