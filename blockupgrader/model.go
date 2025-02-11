package blockupgrader

import (
	"fmt"
)

// schemaModel represents the schema for loading block state upgrade data from a JSON file.
type schemaModel struct {
	MaxVersionMajor    int32 `json:"maxVersionMajor"`
	MaxVersionMinor    int32 `json:"maxVersionMinor"`
	MaxVersionPatch    int32 `json:"maxVersionPatch"`
	MaxVersionRevision int32 `json:"maxVersionRevision"`

	RenamedIDs                  map[string]string                    `json:"renamedIds,omitempty"`
	AddedProperties             map[string]map[string]schemaTagModel `json:"addedProperties,omitempty"`
	RemovedProperties           map[string][]string                  `json:"removedProperties,omitempty"`
	RenamedProperties           map[string]map[string]string         `json:"renamedProperties,omitempty"`
	RemappedPropertyValues      map[string]map[string]string         `json:"remappedPropertyValues,omitempty"`
	RemappedPropertyValuesIndex map[string][]schemaValueRemapModel   `json:"remappedPropertyValuesIndex,omitempty"`
	FlattenedProperties         map[string]schemaFlattenInfo         `json:"flattenedProperties,omitempty"`
	RemappedStates              map[string][]schemaBlockRemapModel   `json:"remappedStates,omitempty"`
}

// schemaBlockRemapModel ...
type schemaBlockRemapModel struct {
	OldProperties    map[string]schemaTagModel `json:"oldState"`
	NewName          string                    `json:"newName"`
	NewFlattenedName schemaFlattenInfo         `json:"newFlattenedName"`
	NewProperties    map[string]schemaTagModel `json:"newState"`
	CopiedProperties []string                  `json:"copiedState"`
}

// schemaFlattenInfo ...
type schemaFlattenInfo struct {
	Prefix                string            `json:"prefix"`
	FlattenedProperty     string            `json:"flattenedProperty"`
	Suffix                string            `json:"suffix"`
	FlattenedValueRemaps  map[string]string `json:"flattenedValueRemaps"`
	FlattenedPropertyType string            `json:"flattenedPropertyType"`
}

// schemaTagModel ...
type schemaTagModel struct {
	Byte   *byte   `json:"byte,omitempty"`
	Int    *int32  `json:"int,omitempty"`
	String *string `json:"string,omitempty"`
}

// schemaValueRemapModel ...
type schemaValueRemapModel struct {
	Old schemaTagModel `json:"old"`
	New schemaTagModel `json:"new"`
}

// parseSchemaModel ...
func parseSchemaModel(m schemaModel) (schema, error) {
	s := schema{
		id: (m.MaxVersionMajor << 24) |
			(m.MaxVersionMinor << 16) |
			(m.MaxVersionPatch << 8) |
			m.MaxVersionRevision,

		maxVersionMajor:    m.MaxVersionMajor,
		maxVersionMinor:    m.MaxVersionMinor,
		maxVersionPatch:    m.MaxVersionPatch,
		maxVersionRevision: m.MaxVersionRevision,

		renamedIDs:          m.RenamedIDs,
		removedProperties:   m.RemovedProperties,
		renamedProperties:   m.RenamedProperties,
		flattenedProperties: m.FlattenedProperties,

		addedProperties:        make(map[string]map[string]any),
		remappedStates:         make(map[string][]schemaBlockRemap),
		remappedPropertyValues: make(map[string]map[string][]schemaValueRemap),
	}

	for blockName, properties := range m.AddedProperties {
		s.addedProperties[blockName] = make(map[string]any)
		for propName, tag := range properties {
			val, err := parseModelTag(tag)
			if err != nil {
				return schema{}, fmt.Errorf("failed to parse model tag: %v", err)
			}
			s.addedProperties[blockName][propName] = val
		}
	}

	convertedRemappedValuesIndex := make(map[string][]schemaValueRemap)
	for mappingKey, mappingValues := range m.RemappedPropertyValuesIndex {
		for _, oldNew := range mappingValues {
			oldVal, err := parseModelTag(oldNew.Old)
			if err != nil {
				return schema{}, fmt.Errorf("failed to parse model tag: %v", err)
			}
			newVal, err := parseModelTag(oldNew.New)
			if err != nil {
				return schema{}, fmt.Errorf("failed to parse model tag: %v", err)
			}
			convertedRemappedValuesIndex[mappingKey] = append(convertedRemappedValuesIndex[mappingKey], schemaValueRemap{
				old: oldVal,
				new: newVal,
			})
		}
	}

	for blockName, properties := range m.RemappedPropertyValues {
		s.remappedPropertyValues[blockName] = make(map[string][]schemaValueRemap)
		for propName, mappedValuesKey := range properties {
			if _, ok := convertedRemappedValuesIndex[mappedValuesKey]; !ok {
				return schema{}, fmt.Errorf("missing key from schema values index: %v", mappedValuesKey)
			}
			s.remappedPropertyValues[blockName][propName] = convertedRemappedValuesIndex[mappedValuesKey]
		}
	}

	for oldBlockName, remaps := range m.RemappedStates {
		for _, remap := range remaps {
			oldProperties, newProperties := make(map[string]any), make(map[string]any)
			for name, tag := range remap.OldProperties {
				val, err := parseModelTag(tag)
				if err != nil {
					return schema{}, fmt.Errorf("failed to parse model tag: %v", err)
				}
				oldProperties[name] = val
			}
			for name, tag := range remap.NewProperties {
				val, err := parseModelTag(tag)
				if err != nil {
					return schema{}, fmt.Errorf("failed to parse model tag: %v", err)
				}
				newProperties[name] = val
			}
			s.remappedStates[oldBlockName] = append(s.remappedStates[oldBlockName], schemaBlockRemap{
				oldProperties:    oldProperties,
				newName:          remap.NewName,
				newFlattenedName: remap.NewFlattenedName,
				newProperties:    newProperties,
				copiedProperties: remap.CopiedProperties,
			})
		}
	}

	return s, nil
}

// parseModelTag ...
func parseModelTag(tag schemaTagModel) (any, error) {
	if tag.Byte != nil {
		return *tag.Byte, nil
	}
	if tag.Int != nil {
		return *tag.Int, nil
	}
	if tag.String != nil {
		return *tag.String, nil
	}
	return nil, fmt.Errorf("invalid tag: %v", tag)
}
