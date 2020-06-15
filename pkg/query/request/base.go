package request

import (
	"github.com/gsow-swc/coc/pkg/http"
)

var (
	token string
)

// SetToken sets the token to be used on requests sent to Clash of Clans
func SetToken(t string) {
	token = t
}

type request interface {
	getURL() string
}

// get retrieves the requested URL and return the results as a byte array.
func get(r request) ([]byte, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}
	client := http.Client{Headers: headers}
	url := r.getURL()
	body, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}
