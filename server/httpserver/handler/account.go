package handler

import (
	"net/http"

	"github.com/Chad-Glazier/controlward/riot"
)

type AccountDto struct {
	GameName string      `json:"gameName"`
	TagLine  string      `json:"tagLine"`
	Puuid    string      `json:"puuid"`
	Server   riot.Server `json:"server"`
}

// Gets a player's account information.
//
// The request must include "gameName" and "tagLine" path parameters.
func GetAccount(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Mandatory path parameters.
		gameName := r.PathValue("gameName")
		if err := riot.ValidateGameName(gameName); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tagLine := r.PathValue("tagLine")
		if err := riot.ValidateTagLine(tagLine); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		account, err := c.Riot.GetAccount(gameName, tagLine)
		if err != nil {
			switch err {
			case riot.ErrNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case riot.ErrRateLimit:
				http.Error(w, err.Error(), http.StatusTooManyRequests)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		accountRegion, err := c.Riot.GetAccountRegion(account.Puuid)
		if err != nil {
			switch err {
			case riot.ErrRateLimit:
				http.Error(w, err.Error(), http.StatusTooManyRequests)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		accountDto := AccountDto{
			GameName: account.GameName,
			TagLine:  account.TagLine,
			Puuid:    account.Puuid,
			Server:   accountRegion.Server,
		}

		sendJson(w, accountDto)
	}
}
