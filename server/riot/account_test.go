package riot

import "testing"

func TestGetAccount(t *testing.T) {
	client := NewClient()

	_, err := client.GetAccount("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAccountRegion(t *testing.T) {
	client := NewClient()

	account, err := client.GetAccount("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := client.GetAccountRegion(account.Puuid)
	if err != nil {
		t.Fatal(err)
	}

	if acc.Server != ServerNA1 {
		t.Fatalf("expected server %s, got %s", ServerNA1, acc.Server)
	}
}
