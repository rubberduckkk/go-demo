package reflect

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 1, 2}
	equal := reflect.DeepEqual(s1, s2)
	t.Logf("equal %v", equal)
}
