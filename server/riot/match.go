package riot

import (
	"compress/gzip"
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
	MatchRanked   = "ranked"
	MatchNormal   = "normal"
	MatchTourney  = "tourney"
	MatchTutorial = "tutorial"
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

// Represents a League of Legends match.
type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Metadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"` // PUUIDs
}

// Contains end of game information.
type Info struct {
	EndOfGameResult string `json:"endOfGameResult,omitempty"`

	GameCreation       int64 `json:"gameCreation"`
	GameDuration       int64 `json:"gameDuration"`
	GameEndTimestamp   int64 `json:"gameEndTimestamp,omitempty"`
	GameStartTimestamp int64 `json:"gameStartTimestamp,omitempty"`

	GameID      int64  `json:"gameId"`
	GameMode    string `json:"gameMode"`
	GameName    string `json:"gameName"`
	GameType    string `json:"gameType"`
	GameVersion string `json:"gameVersion"`

	MapID      int    `json:"mapId"`
	PlatformID string `json:"platformId"`
	QueueID    int    `json:"queueId"`

	TournamentCode string `json:"tournamentCode,omitempty"`

	Participants []Participant `json:"participants"`
	Teams        []Team        `json:"teams"`
}

type Participant struct {
	Puuid          string `json:"puuid"`
	RiotIDGameName string `json:"riotIdGameName,omitempty"`
	RiotIDTagline  string `json:"riotIdTagline,omitempty"`

	SummonerID   string `json:"summonerId,omitempty"`
	SummonerName string `json:"summonerName,omitempty"`

	ProfileIcon int `json:"profileIcon"`

	ParticipantID      int    `json:"participantId"`
	TeamID             int    `json:"teamId"`
	TeamPosition       string `json:"teamPosition"`
	IndividualPosition string `json:"individualPosition"`

	Role string `json:"role"`
	Lane string `json:"lane"`

	ChampionID        int    `json:"championId"`
	ChampionName      string `json:"championName"`
	ChampionTransform int    `json:"championTransform"`

	ChampLevel      int `json:"champLevel"`
	ChampExperience int `json:"champExperience"`

	Kills   int `json:"kills"`
	Deaths  int `json:"deaths"`
	Assists int `json:"assists"`

	DoubleKills int `json:"doubleKills"`
	TripleKills int `json:"tripleKills"`
	QuadraKills int `json:"quadraKills"`
	PentaKills  int `json:"pentaKills"`
	UnrealKills int `json:"unrealKills"`

	LargestKillingSpree int `json:"largestKillingSpree"`
	LargestMultiKill    int `json:"largestMultiKill"`

	TotalDamageDealt    int `json:"totalDamageDealt"`
	MagicDamageDealt    int `json:"magicDamageDealt"`
	PhysicalDamageDealt int `json:"physicalDamageDealt"`
	TrueDamageDealt     int `json:"trueDamageDealt"`

	TotalDamageDealtToChampions    int `json:"totalDamageDealtToChampions"`
	MagicDamageDealtToChampions    int `json:"magicDamageDealtToChampions"`
	PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
	TrueDamageDealtToChampions     int `json:"trueDamageDealtToChampions"`

	TotalDamageTaken    int `json:"totalDamageTaken"`
	PhysicalDamageTaken int `json:"physicalDamageTaken"`
	MagicDamageTaken    int `json:"magicDamageTaken"`
	TrueDamageTaken     int `json:"trueDamageTaken"`

	DamageSelfMitigated int `json:"damageSelfMitigated"`

	DamageDealtToBuildings  int `json:"damageDealtToBuildings"`
	DamageDealtToObjectives int `json:"damageDealtToObjectives"`
	DamageDealtToTurrets    int `json:"damageDealtToTurrets"`

	GoldEarned int `json:"goldEarned"`
	GoldSpent  int `json:"goldSpent"`

	ItemsPurchased int `json:"itemsPurchased"`

	Item0 int `json:"item0"`
	Item1 int `json:"item1"`
	Item2 int `json:"item2"`
	Item3 int `json:"item3"`
	Item4 int `json:"item4"`
	Item5 int `json:"item5"`
	Item6 int `json:"item6"`

	ConsumablesPurchased int `json:"consumablesPurchased"`

	VisionScore             int `json:"visionScore"`
	VisionWardsBoughtInGame int `json:"visionWardsBoughtInGame"`
	WardsPlaced             int `json:"wardsPlaced"`
	WardsKilled             int `json:"wardsKilled"`
	DetectorWardsPlaced     int `json:"detectorWardsPlaced"`

	BaronKills     int `json:"baronKills"`
	DragonKills    int `json:"dragonKills"`
	TurretKills    int `json:"turretKills"`
	InhibitorKills int `json:"inhibitorKills"`

	FirstBloodKill   bool `json:"firstBloodKill"`
	FirstBloodAssist bool `json:"firstBloodAssist"`
	FirstTowerKill   bool `json:"firstTowerKill"`
	FirstTowerAssist bool `json:"firstTowerAssist"`

	TotalMinionsKilled   int `json:"totalMinionsKilled"`
	NeutralMinionsKilled int `json:"neutralMinionsKilled"`

	Win bool `json:"win"`

	GameEndedInSurrender      bool `json:"gameEndedInSurrender"`
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender"`

	TimeCCingOthers        int `json:"timeCCingOthers"`
	LongestTimeSpentLiving int `json:"longestTimeSpentLiving"`

	Perks Perks `json:"perks"`
}

// Perks contains rune information.
type Perks struct {
	StatPerks PerkStats   `json:"statPerks"`
	Styles    []PerkStyle `json:"styles"`
}

type PerkStats struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStyle struct {
	Description string               `json:"description"`
	Selections  []PerkStyleSelection `json:"selections"`
	Style       int                  `json:"style"`
}

type PerkStyleSelection struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

type Team struct {
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
	Bans       []Ban      `json:"bans"`
	Objectives Objectives `json:"objectives"`
}

type Ban struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type Objectives struct {
	Baron      Objective `json:"baron"`
	Champion   Objective `json:"champion"`
	Dragon     Objective `json:"dragon"`
	Inhibitor  Objective `json:"inhibitor"`
	RiftHerald Objective `json:"riftHerald"`
	Tower      Objective `json:"tower"`
}

type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
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

// Decodes a response body as JSON. If the "Content-Encoding" header indicates
// that it's compressed in a format we recognize, then the body will be
// decompressed accordingly. At the time of writing, the recognized encodings
// are: gzip, x-gzip.
//
// Note: The response body will not be closed by this function, even though it
// will be fully read.
func decodeBody(dst any, resp *http.Response) error {

	var body io.Reader

	switch resp.Header.Get("Content-Encoding") {
	case "gzip", "x-gzip":
		r, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer r.Close()
		body = r
	default:
		body = resp.Body
	}

	decoder := json.NewDecoder(body)
	return decoder.Decode(dst)
}

