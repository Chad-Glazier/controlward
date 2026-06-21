package riot

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	Region Region
	Server Server
	Token  string
	Logger *slog.Logger
}

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

const baseDomain = "api.riotgames.com"

func setToken(cfg *Config, r *http.Request) {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	r.Header.Add("X-Riot-Token", riotToken)
}

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
