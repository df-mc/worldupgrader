package itemupgrader

// Item holds the data that identifies an item. It is implemented by ItemMeta.
type Item interface {
	upgrade() ItemMeta
}

// Upgrade upgrades the given item using the registered item upgrade schemas.
// If an Item has not been changed through several versions, Upgrade
// will simply return the original value. Calling itemupgrader.Upgrade is
// therefore safe regardless of whether the item is already up-to-date or not.
func Upgrade(b Item) ItemMeta {
	return b.upgrade()
}

// ItemMeta holds the name and meta values of an item.
type ItemMeta struct {
	Name string
	Meta int16
}

// upgrade upgrades an ItemMeta to a new ItemMeta, changing its Name and Meta if necessary.
func (item ItemMeta) upgrade() ItemMeta {
	for _, s := range schemas {
		if name, ok := s.RenamedIDs[item.Name]; ok {
			item.Name = name
			continue
		}
		if remappedMetas, ok := s.RemappedMetas[item.Name]; ok {
			if newName, ok := remappedMetas[item.Meta]; ok {
				item.Name = newName
				item.Meta = 0
				continue
			}
		}
	}
	return item
}
