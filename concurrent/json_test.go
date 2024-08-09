package main

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestConcurrentMarshal(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	start := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		start <- struct{}{}
		raw, err := json.Marshal(m)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("marshal completed: %s", string(raw))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-start
		m["d"] = 6
	}()

	wg.Wait()
}
