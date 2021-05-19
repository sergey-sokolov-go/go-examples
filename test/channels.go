package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// redact
	cc := make([]chan bool, 3)
	for i := 0; i < cap(cc); i++ {
		cc[i] = make(chan bool)
	}
	go func() {
		time.Sleep(1 * time.Millisecond)
		cc[0] <- true
	}()
	go func() {
		for _, value := range []int{1, 4, 7} {
			// redact
			<-cc[0]
			fmt.Println(value)
			cc[1] <- true
			// redact
		}

		wg.Done()
	}()
	go func() {
		for _, value := range []int{2, 5, 8} {
			// redact
			<-cc[1]
			fmt.Println(value)
			cc[2] <- true
			// redact
		}

		wg.Done()
	}()

	go func() {
		for _, value := range []int{3, 6, 9} {
			// redact
			<-cc[2]
			fmt.Println(value)
			if value != 9 {
				cc[0] <- true
			}
			// redact
		}
		wg.Done()
	}()

	wg.Wait()
}
