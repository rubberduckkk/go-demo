package main

import (
	"testing"
	"time"
)

func TestDone(t *testing.T) {
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		done <- struct{}{}
		t.Logf("finished done")
	}()

	t.Logf("waiting for done")
	<-done
	t.Logf("recv done")
}
