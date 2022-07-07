package main

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s string
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for _, elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	ts := threadSafeSet{s: "abc"}
	iter := ts.Iter()
	for v := range iter {
		fmt.Printf("%c\n", v)
	}
}
