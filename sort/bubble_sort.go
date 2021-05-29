package sort

import (
	"sort"
)

func BubbleSort(data sort.Interface) {
	flag := true
	for i := 0; i < data.Len()-1 && flag; i++ {
		flag = false
		for j := 0; j < data.Len()-i-1; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
				flag = true
			}
		}
	}
}
