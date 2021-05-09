package main

import (
	"fmt"
	"github.com/xXxRisingTidexXx/parallel-computing-labs/internal/pp"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	semaphore, group := pp.NewSemaphore(10), &sync.WaitGroup{}
	group.Add(60)
	for i := 0; i < 60; i++ {
		go work(i, semaphore, group)
	}
	group.Wait()
	fmt.Printf("elapsed time %s\n", time.Since(now))
}

func work(i int, semaphore *pp.Semaphore, group *sync.WaitGroup) {
	semaphore.Lock()
	time.Sleep(time.Second)
	fmt.Printf("worker-%d\n", i)
	semaphore.Unlock()
	group.Done()
}
