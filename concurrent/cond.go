package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

// https://lukechampine.com/cond.html
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
