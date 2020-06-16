package cmd2

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gsow-swc/coc/pkg/query/request"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	star      = "⭐"
	zerostar  = ""
	onestar   = star
	twostar   = star + star
	threestar = star + star + star
)

// heros represents the level of heros for a given player
type heroes struct {
	bk int // Barbarian King level
	aq int // Archer Queen level
	gw int // Grand Warden level
	rc int // Royal Champion level
}

// getHeros returns the hero levels
func getHeroes(hs []response.Troop) heroes {
	var heroes heroes

	for _, h := range hs {
		if h.Name == "Barbarian King" {
			heroes.bk = h.Level
		} else if h.Name == "Archer Queen" {
			heroes.aq = h.Level
		} else if h.Name == "Grand Warden" {
			heroes.gw = h.Level
		} else if h.Name == "Royal Champion" {
			heroes.rc = h.Level
		}
	}

	return heroes
}

// getClan finds the clan with the given name.  If no clan is found or of more than one clan
// is found, an error is returned.
func getClan(name string) (response.Clan, error) {
	var clan response.Clan
	req := request.Clans{Name: name}
	clans, err := req.Get()
	if err != nil {
		log.Error("failed to get the response")
		return clan, err
	}

	if err != nil {
		fmt.Println(err)
		return clan, err
	}
	if len(clans) == 0 {
		err := fmt.Errorf("no clan matching %s was found", name)
		fmt.Println(err)
		return clan, err
	}
	if len(clans) > 1 {
		err := fmt.Errorf("found %d clans matching %s", len(clans), name)
		fmt.Println(err)
		for _, clan := range clans {
			fmt.Printf("   %-20s %s\n", clan.Name, clan.Tag)
		}
		return clan, err
	}

	return clans[0], nil
}

// getTag gets the clan tag from the `clan` option or, if that is not present, from the `name` option
func getTag(c *cli.Context) (string, error) {
	// Get the tag of the clan
	tag := c.String("clan")
	if tag == "" {
		name := c.String("name")
		if name == "" {
			cli.ShowCommandHelpAndExit(c, "get", -1)
		}
		clan, err := getClan(name)
		if err != nil {
			return tag, err
		}
		tag = clan.Tag
	}

	return tag, nil
}

// getTime parses the time returned as a string into a time object
func getTime(t string) time.Time {
	layout := "20060102T150405.000Z"
	time, _ := time.Parse(layout, t)
	return time
}

// getTimeLeft returns a string representation of the time left in a war
func getTimeLeft(t time.Time) string {
	b := strings.Builder{}

	endsIn := t.Sub(time.Now())
	hours := int(endsIn.Hours())
	minutes := int(endsIn.Minutes()) - (hours * 60)

	if hours >= 24 || (hours == 24 && minutes > 0) {
		b.WriteString("War starts in\033[1m")
		if hours > 24 {
			h := hours - 24
			if h == 1 {
				b.WriteString(" " + strconv.Itoa(hours-24) + " hour")
			} else {
				b.WriteString(" " + strconv.Itoa(hours-24) + " hours")
			}
		}
		if minutes == 1 {
			b.WriteString(" " + strconv.Itoa(minutes) + " minute")
		} else if minutes > 1 {
			b.WriteString(" " + strconv.Itoa(minutes) + " minutes")
		}
		b.WriteString("\033[0m")
	} else if hours > 0 || minutes > 0 {
		b.WriteString("War ends in\033[1m")
		if hours == 1 {
			b.WriteString(" " + strconv.Itoa(hours) + " hour")
		} else if hours > 1 {
			b.WriteString(" " + strconv.Itoa(hours) + " hours")
		}
		if minutes == 1 {
			b.WriteString(" " + strconv.Itoa(minutes) + " minute")
		} else if minutes > 1 {
			b.WriteString(" " + strconv.Itoa(minutes) + " minutes")
		}
		b.WriteString("\033[0m")
	} else {
		b.WriteString("War has ended")
	}

	return b.String()
}

// getStars returns a string representation of the number of stars that have
// been gained via an attack against a given base.
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
