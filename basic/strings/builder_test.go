package strings

import (
	"strings"
	"testing"
)

func buildStr1(sb *strings.Builder) {
	sb.WriteString("build 1")
}

func buildStr2(sb *strings.Builder) {
	sb.WriteString("build 2")
}

func TestBuilderCopy(t *testing.T) {
	var sb strings.Builder
	buildStr1(&sb)
	buildStr2(&sb)
	t.Log(sb.String())
}
