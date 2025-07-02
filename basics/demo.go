package main

import "fmt"

type T struct {
}

func (t T) Run() {
	fmt.Printf("t run\n")
}

type D struct {
	T
}

func (d D) Run() {
	fmt.Printf("d run\n")
}
func main() {
	d := new(D)
	d.T = *new(T)
	d.Run()
}
