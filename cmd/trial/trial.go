package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	appName           = "trial"
	debugFlag         = "debug"
	swaggerAPIPath    = "/apidocs.json"
	swaggerPath       = "/swagger/"
	swaggerUIPathFlag = "swagger-ui-path"
	version           = "1.0.0"
	token             = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjJlMDgzYzg5LTA2ZTQtNGUzMC1hMDVmLTIzNTIyNzViZTZlZSIsImlhdCI6MTU4OTM5MTcyOCwic3ViIjoiZGV2ZWxvcGVyL2ViZGQ2NWYyLTdiNTgtZTEwMi1hMWZhLTg3ZjE0YWU5MGU0YyIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjE3My45NS4xNDMuNDciXSwidHlwZSI6ImNsaWVudCJ9XX0.7eJ58DesW9SLgrLlsyfnZ-KnLCLSdKnI1OeOZ8BLs2Reqa1W8ejSfdIvVKYeKKaGcDmoUT3kSovmIkwqwOfafw"
)

func getClans(name string) error {
	req := request.Clans{Name: name}
	clans, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		return err
	}
	for _, c := range clans {
		fmt.Println("Name:", c.Name, ", Tag:", c.Tag)
	}

	return nil
}

func getClan(tag string) error {
	req := request.Clan{Tag: tag}
	clan, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	bytes := []byte(clan.String())
	y, _ := yaml.JSONToYAML(bytes)
	fmt.Println(string(y))

	return nil
}

func getClanMembers(tag string) error {
	req := request.ClanMembers{Tag: tag}
	members, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	fmt.Println(members)

	return nil
}

func getClanMembersTHLevel(tag string) error {
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
	// Sort by TH level and then by name
	sort.Slice(players, func(i, j int) bool {
		var res bool
		if players[i].TownHallLevel != players[j].TownHallLevel {
			res = players[i].TownHallLevel > players[j].TownHallLevel
		} else {
			res = strings.ToLower(players[i].Name) < strings.ToLower(players[j].Name)
		}
		return res
	})

	space := ""
	fmt.Fprintf(os.Stdout, "\033[1m%-18s%-5s%-5s%-5s%-5s%-5s\033[0m\n", "Name", "TH", "BK", "AQ", "GW", "RC")
	for _, p := range players {
		var bkLevel int
		var aqLevel int
		var gwLevel int
		var rcLevel int
		for _, hero := range p.Heroes {
			if hero.Name == "Barbarian King" {
				bkLevel = hero.Level
			} else if hero.Name == "Archer Queen" {
				aqLevel = hero.Level
			} else if hero.Name == "Grand Warden" {
				gwLevel = hero.Level
			} else if hero.Name == "Royal Champion" {
				rcLevel = hero.Level
			}
		}
		fmt.Fprintf(os.Stdout, "%-18s%2d%3s%2d%3s%2d%3s%2d%3s%2d\n", p.Name, p.TownHallLevel, space, bkLevel, space, aqLevel, space, gwLevel, space, rcLevel)
	}

	return nil
}

func getClanWars(tag string) error {
	req := request.ClanWars{Tag: tag}
	wars, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	fmt.Println(wars)

	return nil
}

func getClanCurrentWar(tag string) error {
	req := request.ClanCurrentWar{Tag: tag}
	war, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	fmt.Println(war)

	return nil
}

func getClanWarLeagueGroup(tag string) error {
	req := request.ClanWarLeagueGroup{Tag: tag}
	war, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	fmt.Println(war)

	return nil
}

func getPlayer(tag string) error {
	req := request.Player{Tag: tag}
	player, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		fmt.Println(err)
		return err
	}
	fmt.Println(player)

	return nil
}

func startTrial(ctx *cli.Context) error {
	request.SetToken(token)

	//getClans("CircleOf.")
	//getClan("#2PY099VU2")
	//getClanMembers("#2PY099VU2")
	getClanMembersTHLevel("#2PY099VU2") // CircleOf.Swords
	//getClanMembersTHLevel("#299YPV8CO") // Dark~Brigade
	//getClanWars("#2PY099VU2")
	//getClanCurrentWar("#2PY099VU2")
	//getClanWarLeagueGroup("#2PY099VU2")
	//getClanWarLeagueWar("<id>")
	//getPlayer("#LUUY8RYQV")

	return nil
}

// main starts the GSoW application that listens for new requests.
func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Flags = []cli.Flag{}

	app.Action = startTrial
	app.Run(os.Args)
}
