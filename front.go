package front

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spirosoik/go-front/front/serializer"
)

const (
	// The amount of requests current API token can perform for the 10 seconds window.
	headerRateLimit = "X-RateLimit-Limit"

	// The amount of requests left for the 10 seconds window.
	headerRateRemaining = "X-RateLimit-Remaining"

	// The amount of seconds before the limit resets.
	headerRateReset = "X-RateLimit-Reset"
)

// Config to initialize the Front Client
type Config struct {
	// APIToken provided for your account
	APIToken string

	// The base usrl of Front API
	BaseURL string
}

// Gateway to send the HTTP requests
type Gateway struct {
	// The Http client to send requests to API
	client *http.Client

	// The base Url of Front API
	baseURL *url.URL

	// The API token for Front account
	apiToken string

	//Reusable service client
	common service

	// The service to get info about the current company
	Company *CompanyService

	// The service to run CRUD to channels
	Channel *ChannelService

	// The service to run CRUD to inboxes
	Inbox *InboxService

	// The service to run get requests to team
	Team *TeamService

	// The service to run CRUD to teammate
	Teammate *TeammateService
}

//Service contract struct
type service struct {
	gateway *Gateway
}

// New Factory function to create a Front API Client
func New(cfg *Config) (*Gateway, error) {
	if cfg == nil {
		return nil, fmt.Errorf("Config is needed")
	}
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("Base URL is required")
	}
	if cfg.APIToken == "" {
		return nil, fmt.Errorf("API token is required")
	}
	u, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, err
	}

	g := &Gateway{
		client:   http.DefaultClient,
		baseURL:  u,
		apiToken: cfg.APIToken,
	}

	g.common.gateway = g
	g.Company = (*CompanyService)(&g.common)
	g.Channel = (*ChannelService)(&g.common)
	g.Inbox = (*InboxService)(&g.common)
	g.Team = (*TeamService)(&g.common)
	g.Teammate = (*TeammateService)(&g.common)

	return g, nil
}

// NewRequest creates an HTTP request to be used in client
func (g *Gateway) newRequest(method, path string, body interface{}) (*http.Request, error) {

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		err := enc.Encode(body)

		if err != nil {
			return nil, err
		}
	}
	endpoint := fmt.Sprintf("%s/%s", g.baseURL, path)

	req, err := http.NewRequest(method, endpoint, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.apiToken))
	return req, nil
}

//Call Http request
func (g *Gateway) call(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := g.client.Do(req)

	if err != nil {
		return nil, err
	}

	if code := resp.StatusCode; 400 <= code && code <= 499 {
		var e *ErrorResponse

		deserialize := serializer.Decode(&e)
		if err := deserialize(resp); err != nil {
			return nil, err
		}
		return nil, e
	}

	deserialize := serializer.Decode(&v)
	if err := deserialize(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
