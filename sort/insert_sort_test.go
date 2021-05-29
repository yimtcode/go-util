package sort

import "testing"

func TestInsertSort(t *testing.T) {
	var arr IntArray = []int{24, 12, 34, 20, 53, 66, 25, 44}
	InsertSort(&arr)
	t.Log(arr)
}