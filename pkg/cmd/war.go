package cmd

import (
	"fmt"

	"github.com/gsow-swc/coc/pkg/query/request"
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

// WarGet gets information about the current war for a clan
func WarGet(c *cli.Context) error {
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
