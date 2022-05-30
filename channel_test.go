package belajar_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
//Untuk mengirim data, kita bisa gunakan kode : channel <- data
//Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
//Jika selesai, jangan lupa untuk menutup channel menggunakan function close()

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		//time.Sleep(5 * time.Second)
		channel <- "Eko Kurniawan 1"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
	fmt.Println(data)

	//time.Sleep(3 * time.Second)
}

// ================================================================================================= //

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan 2"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

// ================================================================================================= //

//Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
//Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
//Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerima data)

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan 3"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	time.Sleep(3 * time.Second)
}

// ================================================================================================= //

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
		channel <- "Khannedy"
		channel <- "Khannedy"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// ================================================================================================= //

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// ================================================================================================= //

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// ================================================================================================= //

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
