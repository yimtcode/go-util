// Package array provide go array and slice tools function
package array

// FindIndex find index
func FindIndex(length int, compare func(index int) bool) int {
	for i := 0; i < length; i++ {
		if compare(i) {
			return i
		}
	}

	return -1
}

// Any exist element
func Any(length int, compare func(index int) bool) bool {
	index := FindIndex(length, compare)
	if index == -1 {
		return false
	}

	return true
}