package main

import "fmt"

func makeSlice(c int) []int {
	return make([]int, 0, c)
}

func main() {
	//makeSlice(-1)
	rangeOverNilSlice(nil)
}

func rangeOverNilSlice(slice []int) {
	//var slice []int
	for _, value := range slice {
		fmt.Println(value)
	}
}
