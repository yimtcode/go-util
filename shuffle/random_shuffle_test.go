package shuffle

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomSwapShuffle(t *testing.T) {
	// 使用最好初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 示例
	var arr IntArray = []int{1, 9, 10, 23, 34}
	RandomSwapShuffle(&arr)
	fmt.Println(arr)
}
