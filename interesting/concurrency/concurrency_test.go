package concurrency

import "testing"

func testCummulativeSum(t *testing.T, f func(...int64) int64) {
	tests := []struct {
		name    string
		numbers []int64
		res     int64
	}{
		{
			name:    "Simple",
			numbers: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			res:     55,
		},
	}
	for _, tc := range tests {
		res := f(tc.numbers...)
		if res != tc.res {
			t.Errorf("Test failed: expected %d, got %d", tc.res, res)
		}
	}
}

func TestCummulativeSum(t *testing.T) {
	testCummulativeSum(t, cummulativeSum)
}

func TestCummulativeSumConcurrent(t *testing.T) {
	testCummulativeSum(t, cummulativeSumConcurrent)
}

func BenchmarkCummulativeSum(b *testing.B) {
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		cummulativeSum(nums...)
	}
}

func BenchmarkCummulativeSumConcurrent(b *testing.B) {
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		cummulativeSumConcurrent(nums...)
	}
}
