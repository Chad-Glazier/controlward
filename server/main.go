package main

import (
	"fmt"
	"os"

	"github.com/Chad-Glazier/controlward/riot"
)

func main() {

	id, err := riot.PlayerId("hellokittieirl", "NA1")
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}

	matchInfo, err := riot.OngoingMatch(id)
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}

	fmt.Println(matchInfo)
}
