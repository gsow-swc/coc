package main

import (
	"fmt"
	"os"

	"github.com/gsow-swc/coc/pkg/cmd"
	"github.com/gsow-swc/coc/pkg/log"
	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/urfave/cli/v2"
)

const (
	appName           = "coc"
	swaggerAPIPath    = "/apidocs.json"
	swaggerPath       = "/swagger/"
	swaggerUIPathFlag = "swagger-ui-path"
	usage             = "Clash of Clans CLI"
)

var (
	// Version is the application version
	Version string
	// Revision is the application revision
	Revision string
	// Build is the build date for the application
	Build string
)

var (
	commands = []*cli.Command{
		{
			Name:        "clan",
			Usage:       "Retrieve information about clans",
			Description: "Retrieves information about clans",
			Subcommands: []*cli.Command{
				{
					Name:        "ls",
					Usage:       "Retrieves a list of clans",
					Description: "Retrieves a list of clans",
					Action:      cmd.ClanList,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "Name of the clan to search for",
						},
						&cli.StringFlag{
							Name:    "frequency",
							Aliases: []string{"f"},
							Usage:   "Frequency the clan wars",
						},
						&cli.IntFlag{
							Name:    "location",
							Aliases: []string{"g"},
							Usage:   "Location identifier of the clan",
						},
						&cli.IntFlag{
							Name:    "minmembers",
							Aliases: []string{"m"},
							Usage:   "Minmum number of clan members",
						},
						&cli.IntFlag{
							Name:    "maxmembers",
							Aliases: []string{"M"},
							Usage:   "Maximum number of clan members",
						},
						&cli.StringSliceFlag{
							Name:    "label",
							Aliases: []string{"L"},
							Usage:   "List of label IDs, one flag for each label",
						},
						&cli.IntFlag{
							Name:    "minpoints",
							Aliases: []string{"p"},
							Usage:   "Minimum amount of clan points",
						},
						&cli.IntFlag{
							Name:    "minlevel",
							Aliases: []string{"l"},
							Usage:   "Minimum clan level",
						},
					},
				},
				{
					Name:        "get",
					Usage:       "Gets details about a clan",
					Description: "Gets details about a clan",
					Action:      cmd.ClanGet,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
				{
					Name:        "members",
					Usage:       "Retrieves a list of members of a clan",
					Description: "Retrieves a list of members of a clan",
					Action:      cmd.ClanMembers,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
			},
		},
		{
			Name:        "war",
			Usage:       "Retrieves information about wars a clan has participated in",
			Description: "Retrieves information about wars a clan has participated in",
			Subcommands: []*cli.Command{
				{
					Name:        "ls",
					Usage:       "Retrieves the list of wars a clan has paricipated in",
					Description: "Retrieves the list of wars a clan has paricipated in",
					Action:      cmd.WarList,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
				{
					Name:        "current",
					Usage:       "Retrieves information about the current war for a clan",
					Description: "Retrieves information about the current war for a clan",
					Action:      cmd.WarCurrent,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
				{
					Name:        "attack",
					Usage:       "Retrieves details about the current war for a clan",
					Description: "Retrieves details about the current war for a clan",
					Action:      cmd.WarAttack,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
				{
					Name:        "defend",
					Usage:       "Retrieves details about the current war for a clan",
					Description: "Retrieves details about the current war for a clan",
					Action:      cmd.WarDefend,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
					},
				},
				{
					Name:        "roster",
					Usage:       "Retrieves the roster in the current war",
					Description: "Retrieves the roster in the current war",
					Action:      cmd.WarRoster,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
						&cli.IntFlag{
							Name:        "round",
							Aliases:     []string{"r"},
							Usage:       "The CWL round",
							Value:       0,
							DefaultText: "the current war",
						},
					},
				},
			},
		},
		{
			Name:        "cwl",
			Usage:       "Retrieve information about a Clan War League",
			Description: "Retrieves information about a Clan War League",
			Subcommands: []*cli.Command{
				{
					Name:        "score",
					Usage:       "Retrieves an overview of the current Clan War League war",
					Description: "Retrieves an overview of the current Clan War League war",
					Action:      cmd.CwlScoreboard,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
						&cli.IntFlag{
							Name:        "round",
							Aliases:     []string{"r"},
							Usage:       "The CWL round",
							Value:       0,
							DefaultText: "the current war",
						},
					},
				},
				{
					Name:        "attack",
					Usage:       "Retrieves information about attacks in the current Clan War League war",
					Description: "Retrieves information about attacks in the current Clan War League war",
					Action:      cmd.CwlAttack,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
						&cli.IntFlag{
							Name:        "round",
							Aliases:     []string{"r"},
							Usage:       "The CWL round",
							Value:       0,
							DefaultText: "the current war",
						},
					},
				},
				{
					Name:        "defend",
					Usage:       "Retrieves information about defenses in the current Clan War League war",
					Description: "Retrieves information about defenses in the current Clan War League war",
					Action:      cmd.CwlDefend,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
						&cli.IntFlag{
							Name:        "round",
							Aliases:     []string{"r"},
							Usage:       "The CWL round",
							Value:       0,
							DefaultText: "the current war",
						},
					},
				},
				{
					Name:        "roster",
					Usage:       "Retrieves the roster in a Clan War League war",
					Description: "Retrieves the roster in a Clan War League war",
					Action:      cmd.CwlRoster,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "clan",
							Aliases: []string{"c"},
							Usage:   "The ID of the clan",
						},
						&cli.StringFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "The name of the clan",
						},
						&cli.IntFlag{
							Name:        "round",
							Aliases:     []string{"r"},
							Usage:       "The CWL round",
							Value:       0,
							DefaultText: "the current war",
						},
					},
				},
			},
		},
		{
			Name:   "labels",
			Usage:  "Retrieve information about labels in Clash of Clans",
			Action: labels,
		},
		{
			Name:   "leagues",
			Usage:  "Retrieve information about leagues in Clash of Clans",
			Action: leagues,
		},
		{
			Name:   "locations",
			Usage:  "Retrieve information about locations in Clash of Clans",
			Action: locations,
		},
		{
			Name:   "players",
			Usage:  "Retrieve information about players in Clash of Clans",
			Action: players,
		},
	}

	// flags are the set of flags supported by the CoC application
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			EnvVars:     []string{"COC_TOKEN"},
			Usage:       "API token to use for authentication with the Clash of Clans REST server (required)",
			DefaultText: " ",
			Required:    true,
		},
		&cli.StringFlag{
			Name:  "log",
			Usage: "Log level to be used when logging messages",
			Value: "Warn",
		},
	}
)

func labels(ctx *cli.Context) error {
	fmt.Printf("Hello lables %q", ctx.Args().Get(0))
	return nil
}

func leagues(ctx *cli.Context) error {
	fmt.Printf("Hello leagues %q", ctx.Args().Get(0))
	return nil
}

func locations(ctx *cli.Context) error {
	fmt.Printf("Hello locations %q", ctx.Args().Get(0))
	return nil
}

func players(ctx *cli.Context) error {
	fmt.Printf("Hello players %q", ctx.Args().Get(0))
	return nil
}

/*
coc clan warlog ls --clan tag
coc clan war get --clan tag
coc warleague group --clan tag
coc warleague war --war tag

coc label ls
   -- type (clan or player)

coc league ls
coc league get --league tag
coc league season ls
coc league season get --league tag --season tag
coc league war ls
coc league war get --league tag

coc location ls
coc location get --location id
coc rank clan ls --location id
coc rank player ls --location id
coc rank clanvs ls --location id
coc rank playervs ls --location id

coc player get --player tag
*/

// main starts the GSoW application that listens for new requests.
func main() {
	app := &cli.App{
		Name:     appName,
		Commands: commands,
		Flags:    flags,
		Usage:    usage,
		Version:  Version + "+" + Revision + Build,
		Before: func(c *cli.Context) error {
			// Initialize the logging
			l := c.String("log")
			logLevel := log.GetLogLevel(l)
			log.InitializeLogger(logLevel)

			// Set the token
			token := c.String("token")
			if token != "" {
				request.SetToken(token)
			}

			return nil
		},
	}

	app.Run(os.Args)
}
