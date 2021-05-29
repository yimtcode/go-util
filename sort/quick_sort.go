package sort

import (
	"sort"
)

// 快速排序
// 注意：快速排序是一种不稳定的排序
func QuickSort(data sort.Interface) {
	qSort(data, 0, data.Len()-1)
}

func qSort(data sort.Interface, low, high int) {
	if low < high {
		pivotIndex := partition(data, low, high)

		qSort(data, low, pivotIndex-1)
		qSort(data, pivotIndex+1, high)
	}
}

// 获取枢轴索引
func partition(data sort.Interface, low, high int) int {
	pivot := low

	for low < high {
		// 从后向前扫描
		for low < high && !data.Less(high, pivot) {
			high--
		}
		if low < high {
			data.Swap(low, high)
			pivot = high
			low++
		}

		// 从前向后扫描
		for low < high && !data.Less(pivot, low) {
			low++
		}
		if low < high {
			data.Swap(high, low)
			pivot = low
			high--
		}

	}

	return pivot
}