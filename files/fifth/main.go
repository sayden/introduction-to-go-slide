package main

import (
	"fmt"
	"time"
)

var delay time.Duration = 5000

func main() {
	dispatcherCh := make(chan int, 100)
	workersCh := make(chan int)

	for i := 0; i < 20; i++ {
		go worker(i, workersCh)
	}

	go dispatcher(dispatcherCh, workersCh)

	iter := 0
	for {
		fmt.Printf("Queue has size %d\n", len(dispatcherCh))
		if len(dispatcherCh) == 100 {
			delay = 100
		} else if len(dispatcherCh) == 0 {
			delay = 5000
		}
		dispatcherCh <- iter
		iter++
		time.Sleep(100 * time.Millisecond)
	}
}

func dispatcher(c, w chan int) {
	for {
		v := <-c
		w <- v
	}
}

func worker(id int, w chan int) {
	for {
		select {
		case number := <-w:
			fmt.Printf("Worker %d got the number %d\n", id, number)
			time.Sleep(delay * time.Millisecond)
		}
	}
} //END
