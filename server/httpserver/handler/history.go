package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"compress/gzip"

	"github.com/Chad-Glazier/controlward/riot"
)

// Gets a player's paginated match history. By default, only ranked games are
// returned.
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
				MatchType: riot.MatchRanked,
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

		w.Header().Add("Content-Type", "application/json")

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Add("Content-Encoding", "gzip")

			compressor := gzip.NewWriter(w)
			defer compressor.Close()

			encoder := json.NewEncoder(compressor)
			encoder.Encode(matches)

			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(matches)
	}
}

// Parses the "startIndex" and "count" query parameters. This will not return
// an error if the values are absent, instead it will return defaults. An error
// will only be returned if one of the values are present but are not in the
// correct format. The error message is suitable for sending in a plaintext
// error response.
func parseStartCount(r *http.Request) (uint64, uint64, error) {
	var startIndex uint64 = 0
	var count uint64 = 10
	var err error

	if r.URL.Query().Has("startIndex") {
		startIndex, err = strconv.ParseUint(
			r.URL.Query().Get("startIndex"),
			10,
			64,
		)
		if err != nil {
			return 0, 0, errors.New("startIndex must be a positive integer")
		}
	}

	if r.URL.Query().Has("count") {
		count, err = strconv.ParseUint(
			r.URL.Query().Get("count"),
			10,
			64,
		)
		if err != nil {
			return 0, 0, errors.New("count must be a positive integer")
		}
	}

	return startIndex, count, nil
}
