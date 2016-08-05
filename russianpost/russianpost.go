package russianpost

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://tracking.russianpost.ru/rtm34"
)

// A Client manages communication with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	// Login to access the API Service Tracking
	login string
	// Password to access the API Service Tracking
	password string
}

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(serviceLogin, servicePassword string) *Client {
	httpClient := http.DefaultClient

	bURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, baseURL: bURL, login: serviceLogin, password: servicePassword}
	return c
}
