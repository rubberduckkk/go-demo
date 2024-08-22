package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func _main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var mutex sync.RWMutex

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i == 5 {
				fmt.Printf("wait lock\n")
				mutex.Lock()
				fmt.Printf("acquired lock\n")
			} else {
				fmt.Printf("wait rlock_%v\n", i)
				mutex.RLock()
				fmt.Printf("acquired rlock_%v\n", i)
			}
			time.Sleep(time.Duration(r.Intn(5)) * time.Second)
			if i == 5 {
				mutex.Unlock()
				fmt.Printf("released lock\n")
			} else if i == 3 {
				// 会导致死锁，因为写锁会等读锁释放，也会 block 后续读锁获取
				fmt.Printf("don't release rlock_%v\n", i)
			} else {
				fmt.Printf("released rlock_%v\n", i)
				mutex.RUnlock()
			}
		}(i)
	}

	wg.Wait()
}
