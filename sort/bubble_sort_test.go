package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var arr IntArray = []int{24, 12, 34, 20, 53, 66, 25, 44}
	BubbleSort(&arr)
	t.Log(arr)
}
