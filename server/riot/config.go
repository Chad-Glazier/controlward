package riot

import (
	"log/slog"
	"os"
)

var riotToken string
var logger = *slog.New(slog.NewTextHandler(os.Stdout, nil))
const baseDomain = "api.riotgames.com"

func init() {
	riotToken = os.Getenv("RIOT_TOKEN")
	if riotToken == "" {
		panic("missing RIOT_TOKEN environment variable")
	}
}
