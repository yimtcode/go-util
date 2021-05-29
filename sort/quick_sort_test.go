package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	var arr IntArray = []int{55, 22, 44, 67, 35, 77, 18, 69}
	QuickSort(&arr)
	t.Log(arr)
}