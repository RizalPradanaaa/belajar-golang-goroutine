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
