package json

import (
	"testing"

	"github.com/goccy/go-json"
)

type A struct {
	B  *B  `json:"B"`
	F3 int `json:"f3,omitempty"`
}

type B struct {
	F1 string `json:"f1"`
	F2 int    `json:"f2"`
}

func TestUnmarshal(t *testing.T) {
	var a A
	err := json.Unmarshal([]byte(`{}`), &a)
	t.Logf("err: %v\n", err)
	t.Logf("a: %+v\n", a)
}
