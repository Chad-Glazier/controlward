package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Gets the account details of a user based on their in-game name and tag line.
func (c *Client) GetAccount(gameName, tagLine string) (*Account, error) {

	req, err := c.RequestWithRegionalUrl(fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		gameName, tagLine,
	))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if c.Logger != nil {
			c.Logger.Error("failed to make request", "err", err.Error())
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

	if c.Logger != nil {
		c.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}

// Gets the active region of a player for a game they play. It is currently
// assumed that the game is always League of Legends, so this function only
// takes a player UUID.
func (c *Client) GetAccountRegion(puuid string) (*AccountRegion, error) {

	req, err := c.RequestWithRegionalUrl(fmt.Sprintf(
		"/riot/account/v1/region/by-game/%s/by-puuid/%s",
		GameLeagueOfLegends, puuid,
	))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if c.Logger != nil {
			c.Logger.Error("failed to make request", "err", err.Error())
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

	if c.Logger != nil {
		c.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}

//
// Data Transfer Objects
//
// These types define the shape of JSON data related to accounts that will be
// received from the Riot API.
//

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
