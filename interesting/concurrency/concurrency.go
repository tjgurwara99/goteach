package concurrency

import (
	"sync"
	"sync/atomic"
)

const GOROUTINES = 4

func cummulativeSumConcurrent(ints ...int64) int64 {
	var v int64
	length := len(ints)
	lastGoroutine := GOROUTINES - 1
	perGoroutine := length / GOROUTINES
	var wg sync.WaitGroup
	wg.Add(GOROUTINES)

	for g := 0; g < GOROUTINES; g++ {
		go func(g int) {
			start := g * perGoroutine
			end := start + perGoroutine

			if g == lastGoroutine {
				end = length
			}

			var lv int64
			for _, n := range ints[start:end] {
				lv += n
			}
			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()
	return v
}
