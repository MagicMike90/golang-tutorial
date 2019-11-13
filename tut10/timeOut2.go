package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
		//  w.Wait() call will make the timeout() function wait for a matching sync.Done()
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration!")
		return
	}
	//   w.Wait() call will make the timeout() function wait for a matching sync.Done()
	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
	// Once the w.Done() call is executed, the previous timeout() function will return.
	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}
