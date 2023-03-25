package main

import (
	"github.com/df-mc/blockupgrader"
	"github.com/kr/pretty"
)

func main() {
	pretty.Println(blockupgrader.Upgrade(blockupgrader.BlockState{
		Name: "minecraft:wool",
		Properties: map[string]any{
			"color": "red",
		},
		Version: 17825806,
	}))
}
