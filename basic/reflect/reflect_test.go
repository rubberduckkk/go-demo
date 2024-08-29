package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestA struct {
}

type TestB struct {
}

func TestDeepEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 1, 2}
	equal := reflect.DeepEqual(s1, s2)
	t.Logf("equal %v", equal)
}

func TestFuncValue(t *testing.T) {
	fn1 := CreateClosureNoInline("1", new(TestA))
	fn2 := CreateClosureNoInline("2", new(TestB))
	rv1 := reflect.ValueOf(fn1)
	rv2 := reflect.ValueOf(fn2)
	fmt.Printf("rv1 noinline pointer=%v\n", rv1.Pointer())
	fmt.Printf("rv2 noinline pointer=%v\n", rv2.Pointer())
	assert.Truef(t, rv1.Pointer() == rv2.Pointer(), "func pointer same")

	fn3 := CreateClosureAllowInline("1", new(TestA))
	fn4 := CreateClosureAllowInline("2", new(TestB))
	rv3 := reflect.ValueOf(fn3)
	rv4 := reflect.ValueOf(fn4)
	fmt.Printf("rv3 noinline pointer=%v\n", rv3.Pointer())
	fmt.Printf("rv4 noinline pointer=%v\n", rv4.Pointer())
	assert.Falsef(t, rv3.Pointer() == rv4.Pointer(), "func pointer same")
}
