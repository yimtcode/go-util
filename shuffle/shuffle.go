package shuffle

// 洗牌函数接口
type Interface interface {
	Len() int
	Swap(i, j int)
}