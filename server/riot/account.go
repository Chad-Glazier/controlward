package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func accountGet(path string) (*http.Response, error) {
	url, err := url.Parse("https://americas." + baseDomain + path)
	if err != nil {
		panic(err)
	}

	req := http.Request{}
	req.Method = http.MethodGet
	req.URL = url
	req.Header = http.Header{}
	req.Header.Add("X-Riot-Token", riotToken)

	return http.DefaultClient.Do(&req)
}

// Represents a player account.
type AccountDto struct {
	// Player Universal Unique Identifier. Exact length of 78 characters.
	// (Encrypted)
	Puuid string `json:"puuid"`
	// In-game display name for the account.
	//
	// This field may be excluded from the response if the account doesn't have
	// a gameName.
	GameName string `json:"gameName"`
	// In-game tag line for the user account.
	//
	// This field may be excluded from the response if the account doesn't have
	// a tagLine.
	TagLine string `json:"tagLine"`
}

// Gets the account details of a user based on their in-game name and tag line.
func GetAccountDto(gameName, tagLine string) (*AccountDto, error) {

	resp, err := accountGet(fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		gameName, tagLine,
	))
	if err != nil {
		logger.Error("failed to make request", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		account := AccountDto{}
		if err := json.Unmarshal(body, &account); err != nil {
			return nil, err
		}
		return &account, nil
	}

	logger.Error("riot: error response", "status", resp.Status)
	return nil, riotError(resp)
}

// Represents a game.
type Game string

const (
	LeagueOfLegends  Game = "lol"
	TeamfightTactics Game = "tft"
)

// Represents the active region of a player for a game they play.
type AccountRegionDto struct {
	// Player Universal Unique Identifier. Exact length of 78 characters.
	// (Encrypted)
	Puuid string `json:"puuid"`
	// Game to lookup active region. Either "tft" (Teamfight Tactics) or "lol"
	// (League of Legends).
	Game Game `json:"game"`
	// Player active region. For example, na1 or euw1.
	Server Server `json:"region"`
}

// Gets the active region of a player for a game they play. It is currently
// assumed that the game is always League of Legends, so this function only
// takes a player UUID.
func GetAccountRegionDto(puuid string) (*AccountRegionDto, error) {

	resp, err := accountGet(fmt.Sprintf(
		"/riot/account/v1/region/by-game/%s/by-puuid/%s",
		LeagueOfLegends, puuid,
	))
	if err != nil {
		logger.Error("failed to make request", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		accountRegion := AccountRegionDto{}
		if err := json.Unmarshal(body, &accountRegion); err != nil {
			return nil, err
		}
		return &accountRegion, nil
	}

	logger.Error("riot: error response", "status", resp.Status)
	return nil, riotError(resp)
}
