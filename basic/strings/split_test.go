package strings

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	res := strings.Split("", ",")
	t.Logf("split empty strings: %v\n", res)
}
