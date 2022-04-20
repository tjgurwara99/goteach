package concurrency

func cummulativeSum(ints ...int64) int64 {
	var v int64
	for _, n := range ints {
		v += n
	}
	return v
}
