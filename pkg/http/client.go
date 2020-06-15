package http

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	defaultHeaders = map[string]string{
		"Accept": "application/json",
	}
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
)

// Client is the HTTP client used to send the request to a server.
type Client struct {
	Headers map[string]string
}

// Get sends a request and receives the response from a server.
func (c *Client) Get(url string) ([]byte, error) {
	const M = "http.Client.Send"
	log.Debug(M, " -->")
	defer log.Debug(M, " <--")

	// Get the http request
	log.Debug("GET url=", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("failed to get the http request")
		return nil, err
	}

	// Set the default headers
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}

	// Add any custom headers
	if c.Headers != nil && len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}
	}

	// Send the request to Clash of Clans and get the response
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("failed to send the request to CoC")
		return nil, err
	}
	defer resp.Body.Close()

	// If an error status code was returned by the server, pass the error back to the invoker
	if resp.StatusCode != 200 {
		log.Error("failed to send the request to CoC, statusCode=", resp.StatusCode, ", status=", resp.Status)
		err := fmt.Errorf("status=%d, reason=%s", resp.StatusCode, resp.Status)
		return nil, err
	}

	// Read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("failed to read the body")
		return nil, err
	}

	// fmt.Println(string(body))
	// All good, so return the response
	return body, nil
}
