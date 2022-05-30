package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

//Masalah Dengan Goroutine
//Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
//Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
//Hal ini bisa menyebabkan masalah yang namanya Race Condition

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x) //hasilnya seharusnta 1jt
}
