package belajargolanggoroutine

import (
	"fmt"
	"strconv"
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


// Membuat Channel
func TestChannel(t *testing.T) {
	// Membuat Channel
	channel := make(chan string)

	// Mengisi channel / mengirim data
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Rizal Pradana"
	}()

	// Menerima data dari channel
	data := <- channel
	fmt.Println(data)
	// Untuk menutup channel
	close(channel)
}


// Channel Sebagai Paramater
func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Rizal Pradana"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}

// Membuat Channel In dan Out
// Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan
// untuk mengirim data, atau hanya dapat digunakan untuk menerima data
func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Rizal Pradana"
}

func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}


// Membuat Buffered Channel
// Secara default channel itu hanya bisa menerima 1 data
// Untuknya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian

func TestBufferedChannel(t *testing.T) {
	// untuk capacity buffered channel 3
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Rizal"
	channel <- "Nawang"
	channel <- "Pradana"

	println(<- channel)
	println(<- channel)
	println(<- channel)
	// Akan terjadi Deadlock karena kapasitas buffer hanya 3
	// println(<- channel)

	time.Sleep(2 * time.Second)
}


// Membuat Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i <= 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		println(data)
	}

	time.Sleep(2 * time.Second)
}


// Membuat Select Multiple Channel
// Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
// Lalu kita ingin mendapatkan data dari semua channel tersebu
func TestMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select{
		case data := <- channel1:
			println("Data dari channel1", data)
			counter++
		case data := <- channel2:
			println("Data dari channel2", data)
			counter++
		default:
			println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}

	time.Sleep(1 * time.Second)
}
