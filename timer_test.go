package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	// Akan Dijalankan 5 detik kemudian
	time := <- timer.C
	fmt.Println(time)
}


// time.After()
func TestTimerAfter(t *testing.T) {
	channel := time.After(3 * time.Second)

	tick := <- channel
	fmt.Println(tick)
}


// time.AfterFunc()
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1 * time.Second, func() {
		fmt.Println("Execution after 1 second")
		group.Done()
	})

	group.Wait()
}
