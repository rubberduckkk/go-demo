package main

import (
	"log"
	"testing"
)

type DemoInterface interface {
	ShowDemo()
}

type demoImpl struct {
	content string
}

func (d *demoImpl) ShowDemo() {
	log.Printf("demo: %s\n", d.content)
}

func ReassignDemo(d DemoInterface) {
	if _, ok := d.(*demoImpl); ok {
		d = &demoImpl{
			content: "new content",
		}
	}
}

func ModifyDemo(d DemoInterface) {
	if dd, ok := d.(*demoImpl); ok {
		dd.content = "new content"
	}
}

func ShowDemo(d DemoInterface) {
	d.ShowDemo()
}

func TestAssign(t *testing.T) {
	d := new(demoImpl)
	d.content = "original content"
	ShowDemo(d) // prints "original content"
	ReassignDemo(d)
	ShowDemo(d) // still prints "original content"
	ModifyDemo(d)
	ShowDemo(d) // prints "new content"
}
