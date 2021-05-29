package pagination

import (
	"reflect"
	"testing"
)

func TestPagination(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	startIndex, endIndex := Pagination(len(arr), 2, 7)
	if !reflect.DeepEqual(arr[startIndex:endIndex], []int{8, 9, 10}) {
		t.Fatal("pagination function error")
	}

	startIndex, endIndex = Pagination(len(arr), 3, 7)
	if !reflect.DeepEqual(arr[startIndex:endIndex], []int{}) {
		t.Fatal("pagination function error")
	}
}
