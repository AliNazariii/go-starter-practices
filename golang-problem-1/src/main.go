package main

import (
	"sync"
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	var wg sync.WaitGroup
	var lock sync.Mutex

	numberOfAccesses = 0

	for _, c := range ch {
		wg.Add(1)
		go func(c chan string) {
			select {
			case c <- message:
				lock.Lock()
				defer lock.Unlock()
				numberOfAccesses++
			case <-time.After(d * time.Second):
				break
			}
			wg.Done()
		}(c)
	}

	wg.Wait()
	return numberOfAccesses
}
