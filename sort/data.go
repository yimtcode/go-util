package sort

// 测试演示用
// 多套排序测试共用一个struct
type IntArray []int

func (i *IntArray) Len() int {
	return len(*i)
}

func (i *IntArray) Swap(a, b int) {
	if a == b {
		return
	}
	(*i)[a], (*i)[b] = (*i)[b], (*i)[a]
}

// 通过些方法控制从小到大排序还是从大到小排序
func (i *IntArray) Less(a, b int) bool {
	return (*i)[a] < (*i)[b]
}