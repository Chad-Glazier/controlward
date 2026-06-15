package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Gets a player's ID from their in-game name and tagline.
func PlayerId(gameName, tagLine string) (string, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s/riot/account/v1/accounts/by-riot-id/%s/%s", 
		riotDomain, gameName, tagLine,
	))
	if err != nil {
		panic(err)
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: url,
		Header: headerWithRiotToken(),
	}
	
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic("error: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return "", ErrPlayerNotFound
	}

	jsonData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic("error reading response body: " + err.Error())
	}

	body := make(map[string]any)
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		panic("error parsing JSON from Riot: " + err.Error())
	}

	puuid, ok := body["puuid"]
	if !ok {
		panic("response was ok but body didn't have puuid")
	}

	return puuid.(string), nil
}

// Gets a list of the most recent match ID's associated with a player.
func PlayerMatches(playerId string, start, count uint) ([]string, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s/lol/match/v5/matches/by-puuid/%s/ids?start=%d&count=%d", 
		riotDomain, playerId, start, count,
	))
	if err != nil {
		panic(err)
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: url,
		Header: headerWithRiotToken(),
	}
	
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic("error: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrPlayerNotFound
	}

	jsonData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic("error reading response body: " + err.Error())
	}

	body := []string{}
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		panic("error parsing JSON from Riot: " + err.Error())
	}

	return body, nil
}

// Gets a full match timeline from a match ID.
func Timeline(matchId string) (map[string]any, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s/lol/match/v5/matches/%s/timeline", 
		riotDomain, matchId,
	))
	if err != nil {
		panic(err)
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: url,
		Header: headerWithRiotToken(),
	}
	
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic("error: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrPlayerNotFound
	}

	jsonData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic("error reading response body: " + err.Error())
	}

	body := make(map[string]any)
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		panic("error parsing JSON from Riot: " + err.Error())
	}

	return body, nil
}

// Gets the information about a match.
func MatchInfo(matchId string) (map[string]any, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s/lol/match/v5/matches/%s", 
		riotDomain, matchId,
	))
	if err != nil {
		panic(err)
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: url,
		Header: headerWithRiotToken(),
	}
	
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic("error: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrPlayerNotFound
	}

	jsonData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic("error reading response body: " + err.Error())
	}

	body := make(map[string]any)
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		panic("error parsing JSON from Riot: " + err.Error())
	}

	return body, nil
}

// Gets the information about a player's current game.
func OngoingMatch(playerId string) (map[string]any, error) {
	url, err := url.Parse(fmt.Sprintf(
		"https://%s/lol/spectator/v5/active-games/by-summoner/%s", 
		riotDomain, playerId,
	))
	if err != nil {
		panic(err)
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: url,
		Header: headerWithRiotToken(),
	}
	
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic("error: " + err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrMatchNotFound
	}

	jsonData, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic("error reading response body: " + err.Error())
	}

	body := make(map[string]any)
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		panic("error parsing JSON from Riot: " + err.Error())
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return body, nil
	case http.StatusNotFound:
		return nil, ErrMatchNotFound
	case http.StatusTooManyRequests:
		return nil, ErrRateLimitExceeded
	default:
		return nil, ErrUnknown
	}
}
