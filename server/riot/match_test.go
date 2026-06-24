package riot

import (
	"testing"
)

func TestGetMatchIds(t *testing.T) {
	c := NewClient()

	account, err := c.GetAccount("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}

	gameIds, err := c.GetMatchIds(account.Puuid, 0, 10, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(gameIds) != 10 {
		t.Fatalf("expected %d game IDs, got %d", 10, len(gameIds))
	}
}

func TestGetMatch(t *testing.T) {
	c := NewClient()

	account, err := c.GetAccount("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}

	matchIds, err := c.GetMatchIds(account.Puuid, 0, 4, &OptionsGetMatchIds{
		MatchType: MatchRankedSolo,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(matchIds) != 4 {
		t.Fatalf("expected %d game IDs, got %d", 4, len(matchIds))
	}

	for _, matchId := range matchIds {
		match, err := c.GetMatch(matchId)
		if err != nil {
			t.Fatal(err)
		}

		for _, participant := range match.Info.Participants {
			if participant.Puuid == account.Puuid {
				t.Logf(
					"Champion: %s | KDA: %d/%d/%d | CS: %d | Win: %v",
					participant.ChampionName,
					participant.Kills,
					participant.Deaths,
					participant.Assists,
					participant.TotalMinionsKilled,
					participant.Win,
				)
			}
		}
	}
}
