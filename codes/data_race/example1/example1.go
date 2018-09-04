// Sample program to show how to create race condition in our programs.

// run:  go run -race example1.go

package main

import (
	"fmt"
	"sync"
)

// counter is the variable incremented by all goroutines.
var counter int

func main() {
	// number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// create two goroutines
	for i := 0; i < 2; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// capture the value of counter.
				value := counter

				// increment the local value of counter.
				value++

				// store the value back into counter.
				counter = value
			}
			wg.Done()
		}()
	}

	// wait the goroutine to finish.
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}
