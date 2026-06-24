/*
This package implements a client for making calls to Riot APIs related to user
accounts and League of Legends. You can make a new client with NewClient, or if
you've implemented a cache, NewClientWithCache. The client is very
straightforward; each method calls a single endpoint on the Riot API. This
makes it a little less convenient to use, but also makes it easy to minimize
external calls. It is also strongly recommended that, for "real" projects, you
go through the work of implementing the Cache interface. The cache is designed
to only store data that is large and will not be changed. For example, matches
are always cached by their ID because a match cannot change after it's ended.
Considering the size of a match and how often you'll want to retrieve them,
this saves a lot of network load and API calls with no risk of having outdated
information.
*/
package riot

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

//
// This file implements a Client type, which is used as the object on which
// requests are made. The purpose of having a client struct is to just store
// some repeated configuration between requests.
//

// A configuration for that helps define requests to the Riot API.
type Client struct {
	Region Region
	Token  string
	Logger *slog.Logger
	Cache  Cache
}

// Creates a client with the default configuration.
func NewClient() *Client {
	return &Client{
		Region: RegionAmericas,
		Token:  riotToken,
		Logger: slog.Default(),
	}
}

// Creates a new client with a cache. The cache is used to store certain
// results from Riot that are cache-friendly (i.e., that don't need to be live)
// in order to minimize latency and network load. The cache should be thread-
// safe and declared in the global scope.
func NewClientWithCache(cache Cache) *Client {
	return &Client{
		Region: RegionAmericas,
		Token:  riotToken,
		Logger: slog.Default(),
		Cache:  cache,
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
func (c *Client) RequestWithServerUrl(server Server, path string) (*http.Request, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s.%s%s",
		server,
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
func (c *Client) Ping(server Server) error {
	req, err := c.RequestWithServerUrl(server, "/lol/status/v4/platform-data")
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
