package generic

import (
	"fmt"
	"reflect"
	"testing"
)

func ReAny() any {
	i := 100
	return &i
}

func ReType[T any]() *T {
	re := ReAny()
	i := 5
	b := reflect.TypeOf(re) == reflect.TypeOf(&i)
	fmt.Printf("same type: %v\n", b)
	return re.(*T)
}

func TestReType(t *testing.T) {
	r := ReType[int]()
	fmt.Println(r)
}
