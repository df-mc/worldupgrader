package blockupgrader

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
		file, err := schemasFS.Open("schemas/" + f.Name())
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
	var m schemaModel
	err := json.NewDecoder(r).Decode(&m)
	if err != nil {
		return err
	}
	s, err := parseSchemaModel(m)
	if err != nil {
		return err
	}
	schemas = append(schemas, s)
	return nil
}
