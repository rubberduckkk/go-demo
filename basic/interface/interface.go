package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func process(shape interface{}) {
	if shape == nil {
		return
	}

	switch shape.(type) {
	case *Rectangle:
		fmt.Println("Rectangle Area is:", shape.(*Rectangle).Area())
	default:
		fmt.Printf("unknown type: %T\n", shape)
	}
}

func main() {
	var r *Rectangle
	process(r)
}
