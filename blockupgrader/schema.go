package blockupgrader

// schema ...
type schema struct {
	id int32

	maxVersionMajor    int32
	maxVersionMinor    int32
	maxVersionPatch    int32
	maxVersionRevision int32

	renamedIDs             map[string]string
	addedProperties        map[string]map[string]any
	removedProperties      map[string][]string
	renamedProperties      map[string]map[string]string
	remappedPropertyValues map[string]map[string][]schemaValueRemap
	remappedStates         map[string][]schemaBlockRemap
}

// schemaBlockRemap ...
type schemaBlockRemap struct {
	oldProperties    map[string]any
	newName          string
	newProperties    map[string]any
	copiedProperties []string
}

// schemaValueRemap ...
type schemaValueRemap struct {
	old any
	new any
}

// applyPropertyAdded ...
func (s schema) applyPropertyAdded(oldName string, properties map[string]any) (modified bool) {
	if props, ok := s.addedProperties[oldName]; ok {
		for propName, val := range props {
			if _, ok := properties[propName]; !ok {
				properties[propName] = val
				modified = true
			}
		}
	}
	return modified
}

// applyPropertyRemoved ...
func (s schema) applyPropertyRemoved(oldName string, properties map[string]any) (modified bool) {
	if props, ok := s.removedProperties[oldName]; ok {
		for _, propName := range props {
			if _, ok := properties[propName]; ok {
				delete(properties, propName)
				modified = true
			}
		}
	}
	return modified
}

// applyPropertyRenamedOrValueChanged ...
func (s schema) applyPropertyRenamedOrValueChanged(oldName string, properties map[string]any) (modified bool) {
	if props, ok := s.renamedProperties[oldName]; ok {
		for oldPropName, newPropName := range props {
			if oldVal, ok := properties[oldPropName]; ok {
				delete(properties, oldPropName)
				modified = true

				// If a value remap is needed, we need to do it here, since we won't be able to locate the property
				// after it's been renamed - value remaps are always indexed by old property name for the sake of
				// being able to do changes in any order.
				properties[newPropName] = s.locateNewPropertyValue(oldName, oldPropName, oldVal)
			}
		}
	}
	return modified
}

// applyPropertyValueChanged ...
func (s schema) applyPropertyValueChanged(oldName string, properties map[string]any) (modified bool) {
	if remapped, ok := s.remappedPropertyValues[oldName]; ok {
		for oldPropName := range remapped {
			if oldVal, ok := properties[oldPropName]; ok {
				if newVal := s.locateNewPropertyValue(oldName, oldPropName, oldVal); newVal != oldVal {
					properties[oldPropName] = newVal
					modified = true
				}
			}
		}
	}
	return modified
}

// locateNewPropertyValue ...
func (s schema) locateNewPropertyValue(oldName string, oldPropName string, oldVal any) any {
	if remapped, ok := s.remappedPropertyValues[oldName]; ok {
		if remap, ok := remapped[oldPropName]; ok {
			for _, pair := range remap {
				if pair.old == oldVal {
					return pair.new
				}
			}
		}
	}
	return oldVal
}
