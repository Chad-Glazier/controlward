package main

import (
	"github.com/Chad-Glazier/controlward/cmd"
	"github.com/Chad-Glazier/controlward/riot"
)

func main() {
	println("in game: ", riot.InGame("hello kittie irl", "NA1") == nil)

	cmd.Execute()
}
