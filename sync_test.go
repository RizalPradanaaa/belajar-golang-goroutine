package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
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


// sync.Pool
func TestPool(t *testing.T) {
	var pool =  sync.Pool{
		// Membuat Data Pool Otomatis
		New: func () interface{} {
			return "New"
		},
	}

	// Menambahkan data pada pool
	pool.Put("Rizal")
	pool.Put("Nawang")
	pool.Put("Pradana")

	for i := 1; i <= 10; i++ {
		go func() {
			// Mengambil data pool
			data := pool.Get()
			fmt.Println(data)
			// Menambahkan lagi
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
}


// sync.Map

func TestMap(t *testing.T) {
	var data sync.Map
	var AddData = func(value int)  {
		data.Store(value, value)
	}

	for i := 1; i <= 100; i++ {
		go AddData(i)
	}

	time.Sleep(3 * time.Second)
	data.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}
