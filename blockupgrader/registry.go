package blockupgrader

import (
	"embed"
	"encoding/json"
	"github.com/df-mc/worldupgrader"
	"strings"
)

var (
	//go:embed schemas/*.json
	schemasFS embed.FS
	// schemas is a list of all registered block state upgrade schemas.
	schemas []schema
)

// init ...
func init() {
	files, err := schemasFS.ReadDir("schemas")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}
		buf, err := schemasFS.ReadFile("schemas/" + f.Name())
		if err != nil {
			panic(err)
		}
		var m blockupgrader.schemaModel
		if err = json.Unmarshal(buf, &m); err != nil {
			panic(err)
		}
		s, err := blockupgrader.parseSchemaModel(m)
		if err != nil {
			panic(err)
		}
		schemas = append(schemas, s)
	}
}
