package tesla

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func concurrentPrint(start, end int, numGoroutines int) {
	var val atomic.Int32
	val.Store(int32(start - 1))

	var flag int
	var flagMutex sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		idx := i
		go func(idx int) {
			defer wg.Done()

			for {
				flagMutex.Lock()
				if flag == idx {
					flag++
					flag %= numGoroutines
					newVal := val.Add(1)
					if newVal > int32(end) {
						flagMutex.Unlock()
						return
					}
					fmt.Printf("%v\n", newVal)
				}
				flagMutex.Unlock()
			}
		}(idx)
	}

	wg.Wait()
}

func concurrentPrint2(start, end int, numGoroutines int) {
	var wg sync.WaitGroup
	chs := make([]chan int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		ch := make(chan int)
		chs[i] = ch
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				fmt.Printf("%v\n", val)
			}
		}(ch)
	}

	var flag int
	for i := start; i <= end; i++ {
		flag %= len(chs)
		chs[flag] <- i
		flag++
	}

	for i := 0; i < len(chs); i++ {
		close(chs[i])
	}
	wg.Wait()
}
