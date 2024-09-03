package main

import (
	"fmt"
	"log"
	"reflect"
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

	log.Printf("reflect is nil=%v\n", reflect.ValueOf(shape).IsNil())

	switch shape.(type) {
	case *Rectangle:
		fmt.Println("Rectangle Area is:", shape.(*Rectangle).Area())
	default:
		fmt.Printf("unknown type: %T\n", shape)
	}
}

func main() {
	// will panic
	r := new(Rectangle)
	r = nil
	process(r)

	// will not panic
	process(nil)
}
