package main

import (
	"log"
	"testing"
)

type Demo interface {
	Add(delta int) int
}

type demo1 struct {
	val int
}

func (d *demo1) Add(delta int) int {
	d.val += delta
	return d.val
}

type demo2 struct {
	val int
}

func (d *demo2) Add(delta int) int {
	d.val -= delta
	return d.val
}

func CheckDemo(d Demo) {
	if d == nil {
		log.Println("nil!")
		return
	}
	switch v := d.(type) {
	case *demo1:
		log.Printf("d is *demo1, is_nil=%v\n", v == nil)
	case *demo2:
		log.Printf("d is *demo2\n")
	default:
		log.Printf("d type is: %T\n", v)
	}
}

func TestCheckDemo(t *testing.T) {
	var d *demo1
	CheckDemo(d)
}
