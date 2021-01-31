package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("main: counter = %d\n", counter)
}
