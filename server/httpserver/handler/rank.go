package handler

import (
	"fmt"
	"net/http"

	"github.com/Chad-Glazier/controlward/riot"
)

type RankDto struct {
	Tier         string    `json:"tier"`
	Rank         string    `json:"rank"`
	Puuid        string    `json:"puuid"`
	LeaguePoints int       `json:"leaguePoints"`
	Wins         int       `json:"wins"`
	Losses       int       `json:"losses"`
}

// Gets a player's ranked information.
//
// The request must include "puuid" and "server" path parameters. The server
// parameter should refer to a League of Legends server, like "na1" or "kr".
func GetRank(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Mandatory path parameters.
		puuid := r.PathValue("puuid")
		if err := riot.ValidatePuuid(puuid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		server := r.PathValue("server")
		if err := riot.ValidateServer(server); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		entries, err := c.Riot.GetLeagueEntries(riot.Server(server), puuid)
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

		var soloEntry *riot.LeagueEntry
		for _, entry := range entries {
			if entry.QueueType == riot.QueueRankedSolo {
				soloEntry = &entry
				break
			}
		}
		if soloEntry == nil {
			http.Error(
				w, 
				fmt.Sprintf(
					"no solo queue entry found for %s on %s",
					puuid, server,
				), 
				http.StatusNotFound,
			)
			return
		}

		rank := RankDto{
			Tier: soloEntry.Tier,
			Rank: soloEntry.Rank,
			Puuid: soloEntry.Puuid,
			LeaguePoints: soloEntry.LeaguePoints,
			Wins: soloEntry.Wins,
			Losses: soloEntry.Losses,
		}

		sendJson(w, rank)
	}
}
