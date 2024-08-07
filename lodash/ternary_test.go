package lodash

import (
	"testing"

	"github.com/samber/lo"
)

func TestTernary(t *testing.T) {
	s := []int{1}
	lo.Ternary(len(s) > 1, s[1:2], s[:1])
}
