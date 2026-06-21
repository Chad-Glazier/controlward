package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Represents a player account.
type Account struct {
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
func GetAccount(cfg *Config, gameName, tagLine string) (*Account, error) {

	req := &http.Request{}
	setToken(cfg, req)
	setRegionalUrl(cfg, req, fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		gameName, tagLine,
	))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if cfg.Logger != nil {
			cfg.Logger.Error("failed to make request", "err", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		account := &Account{}
		if err := json.Unmarshal(body, account); err != nil {
			return nil, err
		}
		return account, nil
	}

	if cfg.Logger != nil {
		cfg.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}

// Represents the active region of a player for a game they play.
type AccountRegion struct {
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
func GetAccountRegion(cfg *Config, puuid string) (*AccountRegion, error) {

	req := &http.Request{}
	setToken(cfg, req)
	setRegionalUrl(cfg, req, fmt.Sprintf(
		"/riot/account/v1/region/by-game/%s/by-puuid/%s",
		LeagueOfLegends, puuid,
	))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if cfg.Logger != nil {
			cfg.Logger.Error("failed to make request", "err", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		accountRegion := &AccountRegion{}
		if err := json.Unmarshal(body, accountRegion); err != nil {
			return nil, err
		}
		return accountRegion, nil
	}

	if cfg.Logger != nil {
		cfg.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}
