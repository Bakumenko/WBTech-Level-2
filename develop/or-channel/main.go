package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		var group sync.WaitGroup
		c := make(chan interface{})

		for i := range channels {
			start := time.Now()
			group.Add(1)
			go func(channel <-chan interface{}) {
				for range channel {
				}
				fmt.Printf("close channel after %v\n", time.Since(start))
				group.Done()
			}(channels[i])
		}

		group.Wait()
		close(c)

		return c
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
