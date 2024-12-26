package task1

import (
	"fmt"
	"time"
)

func RunTask1() {
	ch := make(chan int)

	go count(ch)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(time.Second)
}

func count(ch <-chan int) {
	for num := range ch {
		fmt.Println(num * num)
	}
}
