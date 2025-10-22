package ctrl

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
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

func TestDeferLog(t *testing.T) {
	now := time.Now()
	logrus.Infof("defer log start")
	defer logrus.WithField("wrong_cost", time.Since(now)).Infof("defer log end")
	defer func() {
		logrus.WithField("real_cost", time.Since(now)).Infof("defer log end")
	}()
	time.Sleep(time.Second * 5)
}

func TestDeferWithParameters(t *testing.T) {
	now := time.Now()
	logrus.Infof("defer log start")
	defer func(startTime time.Time) {
		logrus.WithField("cost", time.Since(startTime)).Infof("defer log end")
	}(now)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}
