package blockupgrader

// BlockState ...
type BlockState struct {
	Name       string         `nbt:"name"`
	Properties map[string]any `nbt:"states"`
	Version    int32          `nbt:"version"`
}

// Upgrade upgrades the given block state using the registered block state upgrade schemas.
func Upgrade(state BlockState) BlockState {
	version := state.Version
	for _, s := range schemas {
		resVersion := s.id
		if version > resVersion {
			continue
		}

		oldName := state.Name
		oldProperties := state.Properties
		if _, ok := s.remappedStates[oldName]; ok {
			var nextSchema bool
			for _, remap := range s.remappedStates[oldName] {
				if len(oldProperties) != len(remap.oldProperties) {
					continue
				}

				var nextState bool
				for k, v := range oldProperties {
					if _, ok := remap.oldProperties[k]; !ok {
						nextState = true
						break
					}
					if remap.oldProperties[k] != v {
						nextState = true
						break
					}
				}
				if nextState {
					continue
				}

				state, nextSchema = BlockState{
					Name:       remap.newName,
					Properties: remap.newProperties,
					Version:    resVersion,
				}, true
				break
			}
			if nextSchema {
				continue
			}
		}

		properties := state.Properties
		propertyAdded := s.applyPropertyAdded(oldName, properties)
		propertyRemoved := s.applyPropertyRemoved(oldName, properties)
		propertyRenamedOrValueChanged := s.applyPropertyRenamedOrValueChanged(oldName, properties)
		propertyValueChanged := s.applyPropertyValueChanged(oldName, properties)

		name, nameRenamed := s.renamedIDs[oldName]
		if nameRenamed || propertyAdded || propertyRemoved || propertyRenamedOrValueChanged || propertyValueChanged {
			state = BlockState{
				Name:       name,
				Properties: properties,
				Version:    resVersion,
			}
		}
	}
	return state
}
