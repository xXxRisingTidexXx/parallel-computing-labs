package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter uint64
		group   sync.WaitGroup
	)
	for i := 0; i < 50; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Printf("main: counter = %d\n", counter)
}
