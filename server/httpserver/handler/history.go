package handler

import (
	"net/http"

	"github.com/Chad-Glazier/controlward/riot"
)

// Gets a player's paginated match history. By default, only ranked solo games
// are returned.
//
// The request must include "gameName" and "tagLine" path parameters.
func GetHistory(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Optional query parameters.
		startIndex, count, err := parseStartCount(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		matchIds, err := c.Riot.GetMatchIds(
			account.Puuid, startIndex, count,
			&riot.OptionsGetMatchIds{
				MatchType: riot.MatchRankedSolo,
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		matches := make([]riot.Match, count)
		for i, matchId := range matchIds {
			match, err := c.Riot.GetMatch(matchId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			matches[i] = *match
		}

		sendCompressedJson(w, r, matches)
	}
}
