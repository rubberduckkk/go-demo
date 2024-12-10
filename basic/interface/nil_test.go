package main

import (
	"log"
	"reflect"
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

func isNil(b Blaher) bool {
	return b == nil
}

type Blaher interface {
	Blah()
}

type blahImpl struct {
}

func (b *blahImpl) Blah() {
	log.Printf("%v Blah()\n", reflect.TypeOf(b))
}

func TestNil(t *testing.T) {
	var b *blahImpl
	t.Logf("is nil=%v\n", isNil(b))
}

func returnNilError() interface{} {
	return nil
}

func TestInterfaceAssertion(t *testing.T) {
	result := returnNilError()
	err, ok := result.(error)
	t.Logf("err: %v, ok: %v\n", err, ok)
}
