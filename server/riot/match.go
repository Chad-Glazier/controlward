package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Options for the Client.GetMatches function.
type OptionsGetMatchIds struct {
	StartTime uint64
	EndTime   uint64
	MatchType MatchType
}

// Retrieves a list of match IDs for a player. The start index and count are
// used to paginate results.
func (c *Client) GetMatchIds(
	puuid string, startIndex, count uint64,
	opt *OptionsGetMatchIds,
) ([]string, error) {

	req, err := c.RequestWithRegionalUrl(fmt.Sprintf(
		"/lol/match/v5/matches/by-puuid/%s/ids",
		puuid,
	))
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("start", strconv.FormatUint(startIndex, 10))
	query.Add("count", strconv.FormatUint(count, 10))
	if opt != nil {
		if opt.EndTime != 0 {
			query.Add("endTime", strconv.FormatUint(opt.EndTime, 10))
		}
		if opt.StartTime != 0 {
			query.Add("startTime", strconv.FormatUint(opt.StartTime, 10))
		}
		if opt.MatchType != 0 {
			query.Add("queue", strconv.FormatUint(uint64(opt.MatchType), 10))
		}
	}
	req.URL.RawQuery = query.Encode()

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
		ids := make([]string, 0, count)
		if err := json.Unmarshal(body, &ids); err != nil {
			return nil, err
		}
		return ids, nil
	}

	if c.Logger != nil {
		c.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}

// Retrieves highly detailed information about a match and the players
// involved.
func (c *Client) GetMatch(matchId string) (*Match, error) {

	match, _ := c.Cache.LoadMatch(matchId)
	if match != nil {
		c.Logger.Info("cache hit", "matchId", matchId)
		return match, nil
	}

	req, err := c.RequestWithRegionalUrl(fmt.Sprintf(
		"/lol/match/v5/matches/%s",
		matchId,
	))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if c.Logger != nil {
			c.Logger.Error("failed to make request", "err", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		match := &Match{}
		if err := decodeBody(match, resp); err != nil {
			return nil, err
		}

		c.Cache.SaveMatch(matchId, match)
		return match, nil
	}

	if c.Logger != nil {
		c.Logger.Error("riot: error response", "status", resp.Status)
	}
	return nil, riotError(resp)
}

// Gets the details of an ongoing match for the given player.
// func (c *Client) GetCurrentGame(server Server, puuid string) (*)

//
// Data Transfer Objects
//
// These types define the shape of JSON data related to matches that will be
// received from the Riot API.
//

// Represents a League of Legends match.
type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"` // PUUIDs
}

type Info struct {
	EndOfGameResult string `json:"endOfGameResult,omitempty"`

	GameCreation       int64 `json:"gameCreation"`                 // UNIX timestamp
	GameDuration       int64 `json:"gameDuration"`                 // Seconds
	GameEndTimestamp   int64 `json:"gameEndTimestamp,omitempty"`   // UNIX timestamp
	GameStartTimestamp int64 `json:"gameStartTimestamp,omitempty"` // UNIX timestamp

	GameID      int64  `json:"gameId"`
	GameMode    string `json:"gameMode"`
	GameName    string `json:"gameName"`
	GameType    string `json:"gameType"`
	GameVersion string `json:"gameVersion"` // The current League version/patch

	MapID      int       `json:"mapId"`
	PlatformID string    `json:"platformId"` // "NA1", "KR", etc.
	MatchType  MatchType `json:"queueId"`

	TournamentCode string `json:"tournamentCode,omitempty"`

	Participants []Participant `json:"participants"`
	Teams        []Team        `json:"teams"`
}

type Participant struct {
	Puuid          string `json:"puuid"`
	RiotIDGameName string `json:"riotIdGameName"`
	RiotIDTagline  string `json:"riotIdTagline"`

	ChampionLevel int    `json:"champLevel"`
	ChampionId    int    `json:"championId"`
	ChampionName  string `json:"championName"`

	Assists int `json:"assists"`

	BaronKills int `json:"baronKills"`

	PingAllIn    int `json:"allInPings"`    // Yellow crossed swords
	PingAssistMe int `json:"assistMePings"` // Green flag
	PingCommand  int `json:"commandPings"`  // Blue generic ping

	ConsumablesPurchased int `json:"consumablesPurchased"`
	
}
