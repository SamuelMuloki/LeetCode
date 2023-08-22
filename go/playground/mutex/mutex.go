// From https://golangbot.com/mutex/

package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		// its important to pass the addess of the mutex if the mutext is passed by value instead of
		// passing the address, each go routine will have its own copy of the mutex and a race condition will occur
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
