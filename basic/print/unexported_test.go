package print

import (
	"testing"

	"github.com/rubberduckkk/go-demo/basic/print/logpkg"
)

func TestUnexported(t *testing.T) {
	type diffState struct {
		uid  string
		want int
		got  int
	}

	diff := make([]interface{}, 0)
	diff = append(diff, diffState{
		uid:  "123",
		want: 1,
		got:  2,
	})
	logpkg.Print(diff...)
}
