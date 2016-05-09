package main

import (
	"fmt"
	"time"
)

func main() {
	workersCh := make(chan int, 20)
	for i := 0; i < 20; i++ {
		go worker(i, workersCh)
	}

	iter := 0
	for {
		workersCh <- iter
		iter++
		time.Sleep(100 * time.Millisecond)
	}
}

func worker(id int, w chan int) {
	for {
		select {
		case number := <-w:
			fmt.Printf("Worker %d got the number %d\n", id, number)
			time.Sleep(3 * time.Second)
		}
	}
} //END
