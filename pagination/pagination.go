package pagination

// Pagination go slice or array pagination
func Pagination(length, page, size int) (startIndex, endIndex int) {
	start := (page-1)*size + 1
	if start > length {
		startIndex = 0
		endIndex = 0
		return
	}

	endIndex = page * size
	if endIndex > length {
		endIndex = length
	}
	startIndex = start - 1
	return
}
