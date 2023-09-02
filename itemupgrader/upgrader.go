package itemupgrader

import (
	"fmt"
)

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
		name, nameRenamed := s.RenamedIDs[item.Name]
		if !nameRenamed {
			name = item.Name
		}
		meta := item.Meta
		if remappedMetas, ok := s.RemappedMetas[name]; ok {
			if newName, ok := remappedMetas[fmt.Sprintf("%d", meta)]; ok {
				name = newName
				meta = 0
			}
		}
		item = ItemMeta{
			Name: name,
			Meta: meta,
		}
	}
	return item
}
