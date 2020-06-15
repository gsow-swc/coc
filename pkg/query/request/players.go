package request

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/gsow-swc/coc/pkg/config"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
)

// Player contains the parameters to select a given player in SWC
type Player struct {
	Tag string
}

// getURL returns the request URI that may be sent to get a specific player.
func (p *Player) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The player tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/players/")
	sb.WriteString(url.QueryEscape(p.Tag))

	return sb.String()
}

// Get returns the specified player from Clash of Clans.
func (p *Player) Get() (response.Player, error) {
	// Get the player
	body, err := get(p)
	if err != nil {
		return response.Player{}, err
	}

	// Parse into a clan
	var player response.Player
	err = json.Unmarshal(body, &player)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.Player{}, err
	}

	return player, nil
}
