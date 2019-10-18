package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		// The number of goroutines is defined by sync.Add() function.
		waitGroup.Add(1)

		// important to call sync.Add(1) before the go statement
		// in order to prevent any race conditions.
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	// The sync.Wait() call blocks until the counter in the relevant sync.WaitGroup variable is zero,
	// giving your goroutines time to finish.
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
