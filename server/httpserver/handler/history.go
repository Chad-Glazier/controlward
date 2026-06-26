package handler

import (
	"net/http"

	"github.com/Chad-Glazier/controlward/riot"
)

// Gets a player's paginated match history. By default, only ranked solo games
// are returned.
//
// The request must include a "puuid" path parameter.
func GetHistory(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Optional query parameters.
		startIndex, count, err := parseStartCount(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Mandatory path parameters.
		puuid := r.PathValue("puuid")
		if err := riot.ValidatePuuid(puuid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		matchIds, err := c.Riot.GetMatchIds(
			puuid, startIndex, count,
			&riot.OptionsGetMatchIds{
				MatchType: riot.MatchRankedSolo,
			},
		)
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
