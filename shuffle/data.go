package shuffle

// IntArray demo object
type IntArray []int

func (i *IntArray) Len() int {
	return len(*i)
}

func (i *IntArray) Swap(a, b int) {
	if a == b {
		return
	}
	(*i)[a], (*i)[b] = (*i)[b], (*i)[a]
}
