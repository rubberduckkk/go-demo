package main

import (
	//"io"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func test(index int) {
	for i := 0; i < 1000*1000; i++ {
		_ = defaultRand.Intn(20)
	}
	fmt.Println("done", index)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go test(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("exit")
}
