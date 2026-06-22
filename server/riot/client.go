package riot

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

// A configuration for that helps define requests to the Riot API.
type Client struct {
	Region Region
	Server Server
	Token  string
	Logger *slog.Logger
}

// Creates a client with the default configuration.
func NewClient() *Client {
	return &Client{
		Region: RegionAmericas,
		Server: ServerNA1,
		Token:  riotToken,
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

var riotToken string

func init() {
	riotToken = os.Getenv("RIOT_TOKEN")
	if riotToken == "" {
		panic("missing RIOT_TOKEN environment variable")
	}
}

// The base domain name for the Riot API.
const baseDomain = "api.riotgames.com"

// Makes a request object (without sending it) that is configured with the
// Riot token in its header and the base Riot API domain. Paths should include
// a leading "/". The request method is "GET" by default.
func (c *Client) Request(path string) (*http.Request, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s%s",
		baseDomain,
		path,
	))
	if err != nil {
		return nil, err
	}

	req := &http.Request{}
	req.URL = url
	req.Header = http.Header{}
	req.Header.Set("X-Riot-Token", c.Token)
	req.Method = http.MethodGet

	return req, nil
}

// Makes a request object (without sending it) that is configured with the
// Riot token in its header and the regional Riot API domain. Paths should
// include a leading "/". The request method is "GET" by default.
func (c *Client) RequestWithRegionalUrl(path string) (*http.Request, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s.%s%s",
		c.Region,
		baseDomain,
		path,
	))
	if err != nil {
		return nil, err
	}

	req := &http.Request{}
	req.URL = url
	req.Header = http.Header{}
	req.Header.Set("X-Riot-Token", c.Token)
	req.Method = http.MethodGet

	return req, nil
}

// Makes a request object (without sending it) that is configured with the
// Riot token in its header and the game server-specific Riot API domain. Paths
// should include a leading "/". The request method is "GET" by default.
func (c *Client) RequestWithServerUrl(path string) (*http.Request, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s.%s%s",
		c.Server,
		baseDomain,
		path,
	))
	if err != nil {
		return nil, err
	}

	req := &http.Request{}
	req.URL = url
	req.Header = http.Header{}
	req.Header.Set("X-Riot-Token", c.Token)
	req.Method = http.MethodGet

	return req, nil
}

// Pings the Riot server to ensure that the connection is good and the token
// is valid.
func (c *Client) Ping() error {
	req, err := c.RequestWithServerUrl("/lol/status/v4/platform-data")
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return riotError(resp)
}
