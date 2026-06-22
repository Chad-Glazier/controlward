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
