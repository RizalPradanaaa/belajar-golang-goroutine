package belajargolanggoroutine

import (
	"sync"
	"testing"
)

// sync.WaitGroup
func RunAsyncronus(group *sync.WaitGroup)  {
	// Mendai Proses Goroutine Selesai
	defer group.Done()

	// Menandai Prores Goroutine
	group.Add(1)
	println("Hello")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsyncronus(group)
	}

	group.Wait()
	println("Complete")
}
