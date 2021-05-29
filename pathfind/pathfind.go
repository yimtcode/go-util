package pathfind

import (
	"github.com/yimtcode/go-util/array"
)

type MapListInterface interface {
	// 获取指定索引连接的地图
	GetLinks(index int) []interface{}
	// 获取指定地图的数据索引
	GetIndex(v interface{}) int
}

type Path struct {
	index         int   // 当前地图索引
	MapNodes      []int // 已经走过地图索引
	TransferNodes []int // 走过传送阵索引
}

// 返回路线索引
func GetPath(m MapListInterface, start, dest interface{}) *Path {
	startIndex := m.GetIndex(start)
	destIndex := m.GetIndex(dest)
	if startIndex == -1 || destIndex == -1 {
		return nil
	}
	if startIndex == destIndex {
		// 不需要走,有路线为空
		return nil
	}

	startNode := Path{
		index:         startIndex,
		MapNodes:      []int{startIndex},
		TransferNodes: []int{},
	}
	allPath := []Path{startNode}
	already := make([]int, 0)
	return getShortest(m, allPath, destIndex, already)
}

func getShortest(m MapListInterface, allPath []Path, destIndex int, already []int) *Path {
	if len(allPath) == 0 {
		return nil
	}
	newAllPath := make([]Path, 0)

	for _, v := range allPath {
		for subI, subV := range m.GetLinks(v.index) {
			index := m.GetIndex(subV)
			if index == -1 {
				// 没有该地图
				continue
			}

			if array.FindIndex(len(already), func(i int) bool {
				return already[i] == index
			}) != -1 {
				// 该地图已经走过
				continue
			}
			already = append(already, index)

			newPath := Path{}
			newPath.index = index
			newPath.MapNodes = make([]int, len(v.MapNodes), len(v.MapNodes)+1)
			// 过地图索引
			copy(newPath.MapNodes, v.MapNodes)
			newPath.MapNodes = append(newPath.MapNodes, index)
			// 过传送阵索引
			newPath.TransferNodes = make([]int, len(v.TransferNodes), len(v.TransferNodes)+1)
			copy(newPath.TransferNodes, v.TransferNodes)

			newPath.TransferNodes = append(newPath.TransferNodes, subI)
			if index == destIndex {
				return &newPath
			}

			newAllPath = append(newAllPath, newPath)
		}
	}

	return getShortest(m, newAllPath, destIndex, already)
}
