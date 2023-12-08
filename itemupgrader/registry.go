package itemupgrader

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

var (
	//go:embed schemas/*.json
	schemasFS embed.FS
	// schemas is a list of all registered item upgrade schemas.
	schemas []schemaModel
)

// init ...
func init() {
	files, err := schemasFS.ReadDir("remote/id_meta_upgrade_schema")
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
		file, err := schemasFS.Open("remote/id_meta_upgrade_schema/" + f.Name())
		if err != nil {
			panic(fmt.Errorf("failed to open schema: %w", err))
		}
		err = RegisterSchema(file)
		if err != nil {
			panic(fmt.Errorf("failed to register schema: %w", err))
		}
	}
}

// RegisterSchema attempts to decode and parse a schema from the provided file reader. The file must follow the correct
// specification otherwise an error will be returned.
func RegisterSchema(r io.Reader) error {
	var s schemaModel
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		return err
	}
	schemas = append(schemas, s)
	return nil
}
