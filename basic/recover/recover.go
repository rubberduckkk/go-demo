package main

import (
	"fmt"
)

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from", r)
	}
}

func WillRecover() {
	defer Recover()
	panic("Panic will be recovered!")
}

func WontRecover() {
	defer func() {
		Recover()
	}()
	panic("Panic won't be recovered!")
}

func main() {
	WillRecover()
	WontRecover()
}
