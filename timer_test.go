package belajargolanggoroutine

import (
	"fmt"
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
