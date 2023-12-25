package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
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


// sync.Cond
var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCond(value int)  {
	cond.L.Lock()
	group.Wait()
	fmt.Println("Done ", value)
	cond.L.Unlock()
	group.Done()
}

func TestCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		group.Add(1)
		go WaitCond(i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
			// cond.Signal()
		}
	}()

	group.Wait()
}

// Atomic
func TestAtomic(t *testing.T) {
	var group sync.WaitGroup
	var counter int64 = 0
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
