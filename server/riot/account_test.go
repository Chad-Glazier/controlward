package riot

import "testing"

func TestGetAccountDto(t *testing.T) {
	_, err := GetAccountDto("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAccountRegion(t *testing.T) {
	account, err := GetAccountDto("hello kittie irl", "NA1")
	if err != nil {
		t.Fatal(err)
	}

	acc, err := GetAccountRegionDto(account.Puuid)
	if err != nil {
		t.Fatal(err)
	}

	if acc.Server != NA1 {
		t.Fatalf("expected server %s, got %s", NA1, acc.Server)
	}
}
