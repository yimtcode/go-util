package shuffle

import (
	"math/rand"
)

func RandomSwapShuffle(data Interface) {
	rand.Shuffle(data.Len(), func(i, j int) {
		data.Swap(i, j)
	})
}
