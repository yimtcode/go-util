package array

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []int{2, 1, 3, 4}
	index := FindIndex(len(arr), func(index int) bool {
		return arr[index] == 4
	})

	if index != 3 {
		t.Fatal("find index function exception")
	}
}

func TestAny(t *testing.T) {
	arr := []int{2, 1, 3, 4}
	b := Any(len(arr), func(index int) bool {
		return arr[index] == 4
	})
	if !b {
		t.Fatal("any function exception")
	}
}
