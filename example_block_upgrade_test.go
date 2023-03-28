package worldupgrader_test

import (
	"fmt"
	"github.com/df-mc/worldupgrader/blockupgrader"
)

func Example() {
	// BlockState upgrading:
	state := blockupgrader.BlockState{
		Name: "minecraft:wool",
		Properties: map[string]any{
			"color": "red",
		},
		Version: 17825806,
	}
	fmt.Printf("%v\n", blockupgrader.Upgrade(state))
	// Output: {minecraft:red_wool map[] 18040335}
}
