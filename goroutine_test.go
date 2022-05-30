package belajar_goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello Folks")
}

func TestName(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("yasss")

	time.Sleep(1 * time.Second)
}

func TestDulu(t *testing.T) {
	forever := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				forever <- struct{}{}
				return
			default:
				fmt.Println("for loop")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		for {
			for i := 0; i < 5 ; i++ {
				if i == 2 {
					cancel()
				}
			}

		}
	}()

	<-forever
	fmt.Println("finish")
}
