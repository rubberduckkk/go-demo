package main

import (
	"log"
	"sync"
	"testing"
)

func printSlice(i int, s []int) {
	log.Printf("slice[%d] = %v\n", i, s)
}

func TestSlice(T *testing.T) {
	var (
		s        []int
		batchCnt int
		wg       sync.WaitGroup
	)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		log.Printf("len(s)=%d\n", len(s))
		if len(s) >= 5 {
			batchCnt++
			wg.Add(1)
			go func(batch int, s []int) {
				defer wg.Done()
				printSlice(batch, s)
			}(batchCnt, s[:5])
			s = s[5:]
		}
	}
	wg.Wait()
}
