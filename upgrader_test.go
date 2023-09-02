package worldupgrader

import (
	"github.com/df-mc/worldupgrader/blockupgrader"
	"github.com/df-mc/worldupgrader/itemupgrader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockUpgrader(t *testing.T) {
	state := blockupgrader.BlockState{
		Name: "minecraft:wool",
		Properties: map[string]any{
			"color": "red",
		},
		Version: 17825806,
	}
	upgraded := blockupgrader.Upgrade(state)
	assert.Equal(t, "minecraft:red_wool", upgraded.Name)
	assert.Equal(t, map[string]any{}, upgraded.Properties)
	assert.Equal(t, int32(18040335), upgraded.Version)
}

func TestBlockCopiedProperties(t *testing.T) {
	state := blockupgrader.BlockState{
		Name: "minecraft:log",
		Properties: map[string]any{
			"old_log_type": "spruce",
			"pillar_axis":  "y",
		},
		Version: 17825806,
	}
	upgraded := blockupgrader.Upgrade(state)
	assert.Equal(t, "minecraft:spruce_log", upgraded.Name)
	assert.Equal(t, map[string]any{
		"pillar_axis": "y",
	}, upgraded.Properties)
	assert.Equal(t, int32(18042891), upgraded.Version)
}

func TestBlockRemappedPropertyValues(t *testing.T) {
	state := blockupgrader.BlockState{
		Name: "minecraft:big_dripleaf",
		Properties: map[string]any{
			"big_dripleaf_head": byte(1),
			"big_dripleaf_tilt": "none",
			"direction":         int32(2),
		},
		Version: 18090528,
	}
	upgraded := blockupgrader.Upgrade(state)
	assert.Equal(t, "minecraft:big_dripleaf", upgraded.Name)
	assert.Equal(t, map[string]any{
		"big_dripleaf_head":            byte(1),
		"big_dripleaf_tilt":            "none",
		"minecraft:cardinal_direction": "north",
	}, upgraded.Properties)
	assert.Equal(t, int32(18095666), upgraded.Version)
}

func TestItemRenamedID(t *testing.T) {
	item := itemupgrader.ItemMeta{
		Name: "minecraft:record_relic",
	}
	upgraded := itemupgrader.Upgrade(item)
	assert.Equal(t, "minecraft:music_disc_relic", upgraded.Name)
	assert.Equal(t, int16(0), upgraded.Meta)
}

func TestItemRemappedMeta(t *testing.T) {
	item := itemupgrader.ItemMeta{
		Name: "minecraft:concrete",
		Meta: 13,
	}
	upgraded := itemupgrader.Upgrade(item)
	assert.Equal(t, "minecraft:green_concrete", upgraded.Name)
	assert.Equal(t, int16(0), upgraded.Meta)
}
