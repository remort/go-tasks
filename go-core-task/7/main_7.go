package main

import (
	"fmt"
	"sync"
)

func mergeChannels(chNum int, msgNum int) <-chan int {
	chans := make([]chan int, 0, chNum)
	var wg sync.WaitGroup

	for i := range chNum {
		ch := make(chan int)
		chans = append(chans, ch)
		wg.Add(1)
		go func(c chan<- int, id int) {
			defer wg.Done()
			defer close(c)
			for d := range msgNum {
				fmt.Printf("Write %d to ch %d\n", d, id)
				c <- d
			}

		}(ch, i)
	}

	mergeCh := make(chan int)
	for i, ch := range chans {
		wg.Add(1)
		go func(c chan int, id int) {
			defer wg.Done()
			for v := range c {
				fmt.Printf("Send %d from ch %d to merge channel\n", v, id)
				mergeCh <- v
			}
		}(ch, i)

	}

	go func() {
		fmt.Printf("Start awaiting Gorutines \n")
		wg.Wait()
		fmt.Printf("All goroutines finished\n")
		close(mergeCh)
		fmt.Printf("Close merge channel here\n")
	}()

	return mergeCh
}

func main() {
	mergedChannel := mergeChannels(5, 3)

	for v := range mergedChannel {
		fmt.Printf("Read %d from merge channel\n", v)
	}
}
