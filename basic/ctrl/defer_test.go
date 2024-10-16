package ctrl

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func CreateDeferFunc(startTime time.Time, code *int) func() {
	return func() {
		log.Printf("code is %v, cost is %v\n", *code, time.Since(startTime))
	}
}

type Response struct {
	Code int
}

func TestDefer(t *testing.T) {
	r := new(Response)
	defer CreateDeferFunc(time.Now(), &r.Code)()
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	r.Code = 200
}
