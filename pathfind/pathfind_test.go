package pathfind

import (
	"testing"
)

func TestGetPath(t *testing.T) {
	mList := MapList{}
	mList.maps = make([]Map, 0)
	mList.maps = append(mList.maps, Map{
		1,
		[]Link{
			{99, 2},
			{22, 11},
			{66, 12},
			{77, 13},
		},
	})
	mList.maps = append(mList.maps, Map{
		2,
		[]Link{
			{98, 3},
			{21, 21},
			{64, 22},
			{76, 23},
		},
	})
	mList.maps = append(mList.maps, Map{
		3,
		[]Link{
			{97, 4},
			{20, 31},
			{63, 32},
			{75, 33},
		},
	})
	mList.maps = append(mList.maps, Map{
		4,
		[]Link{
			{96, 41},
			{19, 42},
			{62, 43},
		},
	})
	mList.maps = append(mList.maps, Map{32, []Link{}})
	mList.maps = append(mList.maps, Map{62, []Link{}})
	p := GetPath(&mList, 12, 4)
	if p != nil {
		t.Fatal("GetPath test failed")
	}
	p = GetPath(&mList, 1, 32)
	if p == nil {
		t.Fatal("GetPath test failed")
	}
}
