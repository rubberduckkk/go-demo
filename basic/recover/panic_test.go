package main

import (
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestPanic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mockRPCHandler()
	}()
	wg.Wait()
}

func mockRPCHandler() {
	defer func() {
		logrus.Infof("release lock")
	}()
	panicInFunction()
}

func panicInFunction() {
	panic("test")
}
