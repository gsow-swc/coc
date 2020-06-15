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

// Locations lists locations.
type Locations struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of locations that match
// the provided filters.
func (r *Locations) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve eagues
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations")

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

// Get retrieves and returns a list of locations that match the filters.
func (r *Locations) Get() ([]response.Location, error) {
	// Get the leagues
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of leagues
	type respType struct {
		Items []response.Location `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// Location gets information about a specific location.
type Location struct {
	LocationID string // Identifier of the location to retrieve.
}

// getURL returns the request URI that may be sent to get a specific location.
func (r *Location) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the URL to get the requested clan.  The clan tag must be URL encoded to work properly.
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/location/")
	sb.WriteString(url.QueryEscape(r.LocationID))

	return sb.String()
}

// Get returns the requested clan.
func (r *Location) Get() (response.Location, error) {
	// Get the location
	body, err := get(r)
	if err != nil {
		return response.Location{}, err
	}

	// Parse into a location
	var location response.Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.Location{}, err
	}

	return location, nil
}

// LocationClanRankings retrieves clan rankings for a specific location
type LocationClanRankings struct {
	LocationID string // Identifier of the location to retrieve.
	Limit      int    // Limit the number of items returned in the response.
	After      string // Return only items that occur after this marker.
	Before     string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of clan rankings for a
// specific location that match the provided filters.
func (r *LocationClanRankings) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(r.LocationID)
	sb.WriteString("/rankings/clans")

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

// Get retrieves and returns a list of location clan rankings for the given
// location that match the filters.
func (r *LocationClanRankings) Get() (response.LocationClanRanking, error) {
	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return response.LocationClanRanking{}, err
	}

	var lcr response.LocationClanRanking
	err = json.Unmarshal(body, &lcr)
	if err != nil {
		log.Debug("failed to parse the json response")
		return response.LocationClanRanking{}, err
	}

	return lcr, nil
}

// LocationPlayerRankings retrieves player rankings for a specific location
type LocationPlayerRankings struct {
	LocationID string // Identifier of the location to retrieve.
	Limit      int    // Limit the number of items returned in the response.
	After      string // Return only items that occur after this marker.
	Before     string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of player rankings for a
// specific location that match the provided filters.
func (r *LocationPlayerRankings) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(r.LocationID)
	sb.WriteString("/rankings/players")

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

// Get retrieves and returns a list of location player rankings for the given
// location that match the filters.
func (r *LocationPlayerRankings) Get() (response.LocationPlayerRanking, error) {
	var lpr response.LocationPlayerRanking

	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return lpr, err
	}
	err = json.Unmarshal(body, &lpr)
	if err != nil {
		log.Debug("failed to parse the json response")
		return lpr, err
	}

	return lpr, nil
}

// LocationClanVersusRankings retrieves clan versus rankings for a specific location
type LocationClanVersusRankings struct {
	LocationID string // Identifier of the location to retrieve.
	Limit      int    // Limit the number of items returned in the response.
	After      string // Return only items that occur after this marker.
	Before     string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of clan versus rankings for a
// specific location that match the provided filters.
func (r *LocationClanVersusRankings) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(r.LocationID)
	sb.WriteString("/rankings/clans-versus")

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

// Get retrieves and returns a list of location clan versus rankings for the given
// location that match the filters.
func (r *LocationClanVersusRankings) Get() (response.LocationClanVersusRanking, error) {
	var lpr response.LocationClanVersusRanking

	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return lpr, err
	}
	err = json.Unmarshal(body, &lpr)
	if err != nil {
		log.Debug("failed to parse the json response")
		return lpr, err
	}

	return lpr, nil
}

// LocationPlayerVersusRankings retrieves player versus rankings for a specific location
type LocationPlayerVersusRankings struct {
	LocationID string // Identifier of the location to retrieve.
	Limit      int    // Limit the number of items returned in the response.
	After      string // Return only items that occur after this marker.
	Before     string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of player versus rankings for a
// specific location that match the provided filters.
func (r *LocationPlayerVersusRankings) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve league seaons
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(r.LocationID)
	sb.WriteString("/rankings/players-versus")

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

// Get retrieves and returns a list of location player versus rankings for the given
// location that match the filters.
func (r *LocationPlayerVersusRankings) Get() (response.LocationPlayerVersusRanking, error) {
	var lpr response.LocationPlayerVersusRanking

	// Get the league seasons
	body, err := get(r)
	if err != nil {
		return lpr, err
	}
	err = json.Unmarshal(body, &lpr)
	if err != nil {
		log.Debug("failed to parse the json response")
		return lpr, err
	}

	return lpr, nil
}
