package riot

import (
	"net/http"
	"os"
)

var riotToken = ""

const riotDomain = "americas.api.riotgames.com"

func init() {
	riotToken = os.Getenv("RIOT_TOKEN")
	if riotToken == "" {
		panic("missing RIOT_TOKEN environment variable")
	}
}

func headerWithRiotToken() http.Header {
	header := http.Header{}
	header.Add("X-Riot-Token", riotToken)

	return header
}
