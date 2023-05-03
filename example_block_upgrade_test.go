package worldupgrader

import (
	"github.com/df-mc/worldupgrader/blockupgrader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpgrader(t *testing.T) {
	state := blockupgrader.BlockState{
		Name: "minecraft:wool",
		Properties: map[string]any{
			"color": "red",
		},
		Version: 17825806,
	}
	upgraded := blockupgrader.Upgrade(state)
	assert.Equal(t, upgraded.Name, "minecraft:red_wool")
	assert.Equal(t, upgraded.Properties, map[string]any{})
	assert.Equal(t, upgraded.Version, int32(18040335))
}

func TestCopiedProperties(t *testing.T) {
	state := blockupgrader.BlockState{
		Name: "minecraft:log",
		Properties: map[string]any{
			"old_log_type": "spruce",
			"pillar_axis":  "y",
		},
		Version: 17825806,
	}
	upgraded := blockupgrader.Upgrade(state)
	assert.Equal(t, upgraded.Name, "minecraft:spruce_log")
	assert.Equal(t, upgraded.Properties, map[string]any{
		"pillar_axis": "y",
	})
	assert.Equal(t, upgraded.Version, int32(18042891))
}
