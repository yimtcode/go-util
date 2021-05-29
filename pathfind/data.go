package pathfind

type Link struct {
	Transfer int
	MapID    int
}

type Map struct {
	ID    int
	Links []Link
}

type MapList struct {
	maps []Map
}

func (m *MapList) GetLinks(index int) []interface{} {
	ids := make([]interface{}, 0, len(m.maps[index].Links))
	for _, v := range m.maps[index].Links {
		ids = append(ids, v.MapID)
	}
	return ids
}

func (m *MapList) GetIndex(id interface{}) int {
	mapID := id.(int)
	for i, v := range m.maps {
		if v.ID == mapID {
			return i
		}
	}

	return -1
}
