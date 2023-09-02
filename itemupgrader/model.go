package itemupgrader

// schemaModel represents the schema for loading item upgrade data from a JSON file.
type schemaModel struct {
	RenamedIDs    map[string]string            `json:"renamedIds,omitempty"`
	RemappedMetas map[string]map[string]string `json:"remappedMetas,omitempty"`
}
