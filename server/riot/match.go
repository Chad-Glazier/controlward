package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Represents a match type. Either "ranked," "normal," "tourney," or
// "tutorial."
type MatchType string

const (
	MatchRanked = "ranked"
	MatchNormal = "normal"
	MatchTourney = "tourney"
	MatchTutorial = "tutorial"
)

// Options for the Client.GetMatches function.
type OptionsGetMatchIds struct {
	StartTime uint64
	EndTime uint64
	MatchType MatchType
}

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
		if opt.MatchType != "" {
			query.Add("type", string(opt.MatchType))
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

