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


// sync.Once
// Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di
// eksekusi hanya sekali

var counter = 0
func onlyOnce()  {
	counter ++
}

func TestOnce(t *testing.T) {
	var group sync.WaitGroup
	var once sync.Once

	for i := 1; i <= 100; i++ {
		go func() {
			group.Add(1)
			once.Do(onlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	println(counter)
	println("Complete")

}
