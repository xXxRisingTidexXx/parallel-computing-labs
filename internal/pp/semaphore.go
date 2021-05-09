package pp

import (
	"fmt"
)

func NewSemaphore(size int) *Semaphore {
	if size < 1 {
		panic(fmt.Sprintf("sem: semaphore got invalid size, %d < 1", size))
	}
	return &Semaphore{make(chan struct{}, size)}
}

type Semaphore struct {
	hits chan struct{}
}

func (s *Semaphore) Lock() {
	s.hits <- struct{}{}
}

func (s *Semaphore) Unlock() {
	<-s.hits
}
