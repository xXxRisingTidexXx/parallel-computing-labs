package main

import (
	"fmt"
	"sync"
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
				counter++
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Printf("main: counter = %d\n", counter)
}
