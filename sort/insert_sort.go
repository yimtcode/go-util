package sort

import "sort"

func InsertSort(data sort.Interface) {
	for i := 1; i< data.Len(); i++ {
		for j := i-1; j>=0 && data.Less(j+1, j); j-- {
			// 为了兼容sort.Interface接口此处与标准插入排序做了一些修改
			// 注意：未来的新版本中可能放弃兼容sort.Interface接口
			data.Swap(j+1, j)
		}
	}
}