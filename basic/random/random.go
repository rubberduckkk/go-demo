package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randPool = sync.Pool{
		New: func() interface{} {
			return rand.New(rand.NewSource(time.Now().UnixNano()))
		},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < 1000000; i++ {
		_ = random3(1, 100)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func random2(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

func random3(min, max int) int {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Intn(max-min) + min
}
