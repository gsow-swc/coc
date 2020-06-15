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

// Clans is the set of parameters that may be sent to get a list of clans.
type Clans struct {
	Name          string   // Search clans by name.
	WarFrequency  string   // Filter by clan war frequency
	LocationID    int      // Filter by clan location identifier.
	MinMembers    int      // Filter by minimum number of clan members
	MaxMembers    int      // Filter by maximum number of clan members
	MinClanPoints int      // Filter by minimum amount of clan points.
	MinClanLevel  int      // Filter by minimum clan level.
	Limit         int      // Limit the number of items returned in the response.
	After         string   // Return only items that occur after this marker.
	Before        string   // Return only items that occur before this marker.
	LabelIDs      []string // List of label IDs to use for filtering results.
}

// getURL returns the request URI to be sent to get a list of clans that match
// the provided filters.
func (r *Clans) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve clans
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans")

	firstFilter := true
	if r.Name != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("name=")
		sb.WriteString(url.QueryEscape(r.Name))
		firstFilter = false
	}
	if r.WarFrequency != "" {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("warFrequency=")
		sb.WriteString(r.WarFrequency)
		firstFilter = false
	}
	if r.LocationID > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("locationId=")
		sb.WriteString(strconv.Itoa(r.LocationID))
		firstFilter = false
	}
	if r.MinMembers > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("minMembers=")
		sb.WriteString(strconv.Itoa(r.MinMembers))
		firstFilter = false
	}
	if r.MaxMembers > 0 {
		if !firstFilter {
			sb.WriteString("&")
		}
		sb.WriteString("maxMembers=")
		sb.WriteString(strconv.Itoa(r.MaxMembers))
		firstFilter = false
	}
	if r.MinClanPoints > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("minClanPoints=")
		sb.WriteString(strconv.Itoa(r.MinClanPoints))
		firstFilter = false
	}
	if r.MinClanLevel > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("minClanLevel=")
		sb.WriteString(strconv.Itoa(r.MinClanLevel))
		firstFilter = false
	}
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
	if r.LabelIDs != nil && len(r.LabelIDs) > 0 {
		if firstFilter {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString("labelIds=")
		for i, l := range r.LabelIDs {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(l)
		}
		firstFilter = false
	}

	// Return the full URL
	return sb.String()
}

// Get retrieves and returns a list of clans that match the filters.
func (r *Clans) Get() ([]response.Clan, error) {
	// Get the clans
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Items []response.Clan `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// Clan is the parameters that may be sent to get a specific clan.
type Clan struct {
	Tag string // Tag of the clan.
}

// getURL returns the request URI that may be sent to get a specific clan.
func (r *Clan) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(url.QueryEscape(r.Tag))

	return sb.String()
}

// Get returns the requested clan.
func (r *Clan) Get() (response.Clan, error) {
	// Get the clans
	body, err := get(r)
	if err != nil {
		return response.Clan{}, err
	}

	// Parse into a clan
	var clan response.Clan
	err = json.Unmarshal(body, &clan)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.Clan{}, err
	}

	return clan, nil
}

// ClanMembers are the set of parameters that may be sent to get a list of members of a specific clan.
type ClanMembers struct {
	Tag    string // Tag of the clan.
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI that may be sent to get a list of members of a specific clan.
func (r *ClanMembers) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(url.QueryEscape(r.Tag))
	sb.WriteString("/members")

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

// Get returns the requested clan members.
func (r *ClanMembers) Get() ([]response.ClanMember, error) {
	// Get the clan members
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clan members
	type respType struct {
		Items []response.ClanMember `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// ClanWars is the set of parameters that may be sent to get a clan's clan war log.
type ClanWars struct {
	Tag    string // Tag of the clan.
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI that may be sent to get a clan's clan war log.
func (r *ClanWars) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(url.QueryEscape(r.Tag))
	sb.WriteString("/warlog")

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

// Get returns the requested clan members.
func (r *ClanWars) Get() ([]response.ClanWar, error) {
	// Get the clan members
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clan members
	type respType struct {
		Items []response.ClanWar `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// ClanCurrentWar is the set of parameters that may be used to get a clan's current clan war.
type ClanCurrentWar struct {
	Tag string // Tag of the clan.
}

// getURL returns the request URI that may be sent to get a clan's current clan war.
func (r *ClanCurrentWar) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(url.QueryEscape(r.Tag))
	sb.WriteString("/currentwar")

	return sb.String()
}

// Get retrieves the current clan war for the specified clan.
func (r *ClanCurrentWar) Get() (response.ClanWar, error) {
	// Get the clan war
	body, err := get(r)
	if err != nil {
		return response.ClanWar{}, err
	}

	// Parse into a clan war
	var cw response.ClanWar
	err = json.Unmarshal(body, &cw)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.ClanWar{}, err
	}

	return cw, nil
}

// ClanWarLeagueGroup is the set of parameters that may be used to retrieve information about a clan's current clan war league group.
type ClanWarLeagueGroup struct {
	Tag string // Tag of the clan.
}

// getURL returns the request URI that may be used to retrieve information about a clan's current clan war league group.
func (r *ClanWarLeagueGroup) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(url.QueryEscape(r.Tag))
	sb.WriteString("/currentwar/leaguegroup")

	return sb.String()
}

// Get retrieves the current clan war league group for the specified clan.
func (r *ClanWarLeagueGroup) Get() (response.ClanWarLeagueGroup, error) {
	// Get the clan war
	body, err := get(r)
	if err != nil {
		return response.ClanWarLeagueGroup{}, err
	}

	// Parse into a clan war
	var cw response.ClanWarLeagueGroup
	err = json.Unmarshal(body, &cw)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.ClanWarLeagueGroup{}, err
	}

	return cw, nil
}

// ClanWarLeagueWar is the set of parameters that may be used to retrieve information about an individual clan war league war.
type ClanWarLeagueWar struct {
	Tag string // Tag of the war.
}

// getURL returns the request URI that may be used to retrieve information about an individual clan war league war.
func (r *ClanWarLeagueWar) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clanwarleagues/wars/")
	sb.WriteString(url.QueryEscape(r.Tag))

	return sb.String()
}

// Get retrieves the current clan war league group for the specified clan.
func (r *ClanWarLeagueWar) Get() (response.ClanWarLeagueWar, error) {
	// Get the clan war
	body, err := get(r)
	if err != nil {
		return response.ClanWarLeagueWar{}, err
	}

	// Parse into a clan war
	var cw response.ClanWarLeagueWar
	err = json.Unmarshal(body, &cw)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.ClanWarLeagueWar{}, err
	}

	return cw, nil
}
