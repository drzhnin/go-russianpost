package russianpost

import (
	"io"
	"io/ioutil"
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

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
