package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	fmt.Println("testt")
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 1000000; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}

func TestCoba(t *testing.T)  {
	ktk2 := 1

	ktkNew := ktk2 - ktk2

	ktkNew2 := ktkNew + 11

	var ktk5 int

	if ktk5 > ktkNew2{
		fmt.Println("yes")
	} else {
		fmt.Println("no")
		fmt.Println(ktk5)
	}
}