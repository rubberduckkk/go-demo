package main

func makeSlice(c int) []int {
	return make([]int, 0, c)
}

func main() {
	makeSlice(-1)
}
