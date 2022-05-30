package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNum(num int)  {
	fmt.Println("Display", num)
}

func TestManyGoroutine(t *testing.T) {
	for i:=0; i < 10; i++ {
		DisplayNum(i)
	}

	time.Sleep(1 * time.Second)
}
