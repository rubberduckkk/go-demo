package ctrl

import (
	"log"
	"testing"
)

func CreateDeferFunc(code *int) func() {
	return func() {
		log.Printf("code is %v\n", *code)
	}
}

type Response struct {
	Code int
}

func TestDefer(t *testing.T) {
	r := new(Response)
	defer CreateDeferFunc(&r.Code)()

	r.Code = 200
}
