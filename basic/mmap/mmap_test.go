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
