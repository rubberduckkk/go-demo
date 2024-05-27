package main

type MyStruct struct {
	A int
	B int
	C int
}

type YourStruct struct {
	A int
	B int
	C int
}

func getYourStruct() *YourStruct {
	return nil
}

func getC(y int) int {
	return y
}

func main() {
	y := getYourStruct()
	_ = &MyStruct{
		A: y.A,
		B: y.B,
		C: getC(y.C),
	}
}
