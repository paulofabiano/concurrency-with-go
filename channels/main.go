package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch <- 0
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
