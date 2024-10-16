package ctrl

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func CreateDeferFunc(startTime time.Time, code *int) {
	log.Printf("code is %v, cost is %v\n", *code, time.Since(startTime))
}

type Response struct {
	Code int
}

func TestDefer(t *testing.T) {
	log.Printf("processing request...\n")
	r := new(Response)
	defer CreateDeferFunc(time.Now(), &r.Code)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	r.Code = 200
	log.Printf("finished processing...\n")
}
