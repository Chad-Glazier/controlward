package riot

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

// A configuration for that helps define requests to the Riot API.
type Config struct {
	Region Region
	Server Server
	Token  string
	Logger *slog.Logger
}

// Creates the default configuration.
func NewConfig() *Config {
	return &Config{
		Region: Americas,
		Server: NA1,
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

// Sets the configured Riot API token onto the given request.
func setToken(cfg *Config, r *http.Request) {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	r.Header.Add("X-Riot-Token", riotToken)
}

// Sets the URL of a request with the default Riot API domain.
func setUrl(cfg *Config, r *http.Request, path string) error {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s%s",
		baseDomain,
		path,
	))
	if err != nil {
		return err
	}

	r.URL = url
	return nil
}

// Sets the URL of a request using a domain that matches the configured region.
func setRegionalUrl(cfg *Config, r *http.Request, path string) error {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s.%s%s",
		cfg.Region,
		baseDomain,
		path,
	))
	if err != nil {
		return err
	}

	r.URL = url
	return nil
}

// Sets the URL of a request using a domain that matches the configured game
// server.
func setServerUrl(cfg *Config, r *http.Request, path string) error {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s.%s%s",
		cfg.Server,
		baseDomain,
		path,
	))
	if err != nil {
		return err
	}

	r.URL = url
	return nil
}
