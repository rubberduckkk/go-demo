package reflect

import (
	"fmt"
	"log"
	"reflect"
)

func doSomething(arg string) string {
	return arg + "1"
}

func doOther(arg any) any {
	typ := reflect.TypeOf(arg)
	fmt.Printf("typ: %v\n", typ)
	return arg
}

//go:noinline
func CreateClosureNoInline(arg string, arg2 any) func() {
	return func() {
		arg = doSomething(arg)
		arg2 = doOther(arg2)
		log.Printf("arg: %v, arg2: %v\n", arg, arg2)
	}
}

func CreateClosureAllowInline(arg string, arg2 any) func() {
	return func() {
		arg = doSomething(arg)
		arg2 = doOther(arg2)
		log.Printf("arg: %v, arg2: %v\n", arg, arg2)
	}
}
