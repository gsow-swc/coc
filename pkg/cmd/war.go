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

// WarList gets the list of wars a clan has participated in
func WarList(c *cli.Context) error {
	// Get the tag of the clan
	tag, err := getTag(c)
	if err != nil {
		return err
	}

	// Get the clan wars
	req := request.ClanWars{Tag: tag}
	ws, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}

	warList := wars{}
	for _, w := range ws {
		if w.Opponent.Name != "" {
			wr := warResult{
				OpponentName:    w.Opponent.Name,
				OpponentTag:     w.Opponent.Tag,
				TeamSize:        w.TeamSize,
				Result:          w.Result,
				Stars:           w.Clan.Stars,
				Percent:         w.Clan.DestructionPercentage,
				OpponentStars:   w.Opponent.Stars,
				OpponentPercent: w.Opponent.DestructionPercentage,
			}
			warList.Results = append(warList.Results, wr)
		}
	}

	fmt.Println(warList)

	return nil
}

// WarCurrent gets information about the current war for a clan
func WarCurrent(c *cli.Context) error {
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

	warList := wars{}
	wr := warResult{
		OpponentName:    war.Opponent.Name,
		OpponentTag:     war.Opponent.Tag,
		TeamSize:        war.TeamSize,
		Result:          war.Result,
		Stars:           war.Clan.Stars,
		Percent:         war.Clan.DestructionPercentage,
		OpponentStars:   war.Opponent.Stars,
		OpponentPercent: war.Opponent.DestructionPercentage,
	}
	warList.Results = append(warList.Results, wr)
	fmt.Println(warList)

	return nil
}

// WarAttack gets details about the current war for a clan
func WarAttack(c *cli.Context) error {
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
		CwlWar:  false,
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

// WarDefend gets details about the current war for a clan
func WarDefend(c *cli.Context) error {
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

// WarTargets gets the set of non-cleared targets for the current war
func WarTargets(c *cli.Context) error {
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

	var wt warTargets
	wt = warTargets{ClanName: clan.Name, OpponentName: opponent.Name}
	for i, m := range opponent.Members {
		if m.BestOpponentAttack.Stars < 3 {
			target := warStatusTarget{
				Name:        m.Name,
				MapPosition: i + 1,
				TownHall:    m.TownhallLevel,
			}
			if m.BestOpponentAttack.AttackerTag != "" {
				target.Stars = m.BestOpponentAttack.Stars
				target.DestructionPercentage = m.BestOpponentAttack.DestructionPercentage
			}
			wt.Targets = append(wt.Targets, target)
		}
	}

	fmt.Println(wt)
	return nil
}
