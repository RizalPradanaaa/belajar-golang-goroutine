package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}


// Membuat Banyak Goroutine
func DisplayNumber(number int)  {
	fmt.Println("Display", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
