package main

import (
	"fmt"

	"github.com/Chad-Glazier/controlward/riot"
)

func main() {

	id, err := riot.PlayerId("hellokittieirl", "NA1")
	if err != nil {
		panic(err)
	}

	matchIds, err := riot.PlayerMatchIds(id, 0, 10)
	if err != nil {
		panic(err)
	}

	match, err := riot.MatchInfo(matchIds[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(match)
}
