package request

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gsow-swc/coc/pkg/config"
	"github.com/gsow-swc/coc/pkg/query/response"
	log "github.com/sirupsen/logrus"
)

// ClanLabels lists clan labels
type ClanLabels struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of clans labels that match
// the provided filters.
func (r *ClanLabels) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve clans
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/labels/clans")

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

// Get retrieves and returns a list of clan labels that match the filters.
func (r *ClanLabels) Get() ([]response.Label, error) {
	// Get the clan labels
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clan labels
	type respType struct {
		Items []response.Label `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}

// PlayerLabels lists player labels
type PlayerLabels struct {
	Limit  int    // Limit the number of items returned in the response.
	After  string // Return only items that occur after this marker.
	Before string // Return only items that occur before this marker.
}

// getURL returns the request URI to be sent to get a list of clans labels that match
// the provided filters.
func (r *PlayerLabels) getURL() string {
	var sb strings.Builder
	sb.Grow(100)

	// Get the base URL to retrieve clans
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/labels/clans")

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

// Get retrieves and returns a list of player labels that match the filters.
func (r *PlayerLabels) Get() ([]response.Label, error) {
	// Get the player labels
	body, err := get(r)
	if err != nil {
		return nil, err
	}

	// Parse into an array of player labels
	type respType struct {
		Items []response.Label `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Items, nil
}
