package mmap

import (
	"fmt"
	"testing"
)

func getMap() map[string]interface{} {
	return nil
}

func TestNil(t *testing.T) {
	m := getMap()
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func TestDeleteInFor(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = 100 - i
	}

	for k := range m {
		if k < 50 {
			delete(m, k)
		}
	}
	t.Logf("remain: %v", len(m))
}
