package main

import (
	"fmt"
	"runtime"
)

func bar() {
	_, file, no, ok := runtime.Caller(2)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}
}

func foo() {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}
	bar()
}

func main() {
	foo()
}
