package request

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/gsow-swc/coc/pkg/config"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
)

// Leagues lists leagues
type Leagues struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of leagues that match
// the provided filters.
func (r *Leagues) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve eagues
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues")

	firstFilter := true
	if r.Limit > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("limit=")
		sb.WriteString(strconv.Itoa(r.Limit))
		firstFilter = false
	}
	if r.After != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("after=")
		sb.WriteString(r.After)
		firstFilter = false
	}
	if r.Before != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("before=")
		sb.WriteString(r.Before)
		firstFilter = false
	}

	return sb.String()
}

// Get retrieves and returns a list of leagues that match the filters.
func (r *Leagues) Get() ([]response.League, error) {
	// Get the leagues
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of leagues
	type respType struct {
		Items []response.League `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// League gets information about a specific league
type League struct {
	LeagueID string // Identifier of the league.
}

// getURL returns the request URI that may be sent to get a specific league.
func (r *League) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(url.QueryEscape(r.LeagueID))

	return sb.String()
}

// Get returns the requested clan.
func (r *League) Get() (response.League, error) {
	// Get the league
	body, err := get(r)
	if err != nil {
		return response.League{}, err
	}

	// Parse into a league
	var league response.League
	err = json.Unmarshal(body, &league)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.League{}, err
	}

	return league, nil
}

// LeagueSeasons gets league seasons.  Note that leage season information is
// only available for Legend League
type LeagueSeasons struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of league seasons that match
// the provided filters.
func (r *LeagueSeasons) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve eagues
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues")

	firstFilter := true
	if r.Limit > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("limit=")
		sb.WriteString(strconv.Itoa(r.Limit))
		firstFilter = false
	}
	if r.After != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("after=")
		sb.WriteString(r.After)
		firstFilter = false
	}
	if r.Before != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("before=")
		sb.WriteString(r.Before)
		firstFilter = false
	}

	return sb.String()
}

// Get retrieves and returns a list of league seasons that match the filters.
func (r *LeagueSeasons) Get() ([]response.LeagueSeason, error) {
	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of league seasons
	type respType struct {
		Items []response.LeagueSeason `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// LeagueSeason gets league season rankings. Note that league season information is available only for Legend League.
type LeagueSeason struct {
	LeagueID string // Identifier of the league.
	SeasonID string // Identifier of the season.
	Limit    int    // Limit the number of items returned in the response.
	After    string // Return only items that occur after this marker.
	Before   string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of league seasons that match
// the provided filters.
func (r *LeagueSeason) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(r.LeagueID)
	sb.WriteString("/seasons/")
	sb.WriteString(r.SeasonID)

	firstFilter := true
	if r.Limit > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("limit=")
		sb.WriteString(strconv.Itoa(r.Limit))
		firstFilter = false
	}
	if r.After != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("after=")
		sb.WriteString(r.After)
		firstFilter = false
	}
	if r.Before != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("before=")
		sb.WriteString(r.Before)
		firstFilter = false
	}

	return sb.String()
}

// Get retrieves and returns a list of league seasons that match the filters.
func (r *LeagueSeason) Get() (response.LeagueSeason, error) {
	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return response.LeagueSeason{}, err
	}

	var ls response.LeagueSeason
	err = json.Unmarshal(body, &ls)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.LeagueSeason{}, err
	}

	return ls, nil
}

// WarLeagues lists war leagues
type WarLeagues struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of war leagues that match
// the provided filters.
func (r *WarLeagues) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve war leagues
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/warleagues")

	firstFilter := true
	if r.Limit > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("limit=")
		sb.WriteString(strconv.Itoa(r.Limit))
		firstFilter = false
	}
	if r.After != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("after=")
		sb.WriteString(r.After)
		firstFilter = false
	}
	if r.Before != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("before=")
		sb.WriteString(r.Before)
		firstFilter = false
	}

	return sb.String()
}

// Get retrieves and returns a list of war leagues that match the filters.
func (r *WarLeagues) Get() ([]response.WarLeague, error) {
	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of war leagues
	type respType struct {
		Items []response.WarLeague `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// WarLeague gets war league information.
type WarLeague struct {
	LeagueID string // Identifier of the league.
}

// getURL returns the request URI to be sent to get a specific war league
// the provided filters.
func (r *WarLeague) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/warleagues/")
	sb.WriteString(r.LeagueID)

	return sb.String()
}

// Get retrieves and returns a list of league seasons that match the filters.
func (r *WarLeague) Get() (response.WarLeague, error) {
	// Get the war league
	body, err := get(r)
	if err != nil {
		return response.WarLeague{}, err
	}

	var wl response.WarLeague
	err = json.Unmarshal(body, &wl)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.WarLeague{}, err
	}

	return wl, nil
}
