package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

// The first part creates three global variable
var (
	size = 10
	// getting new client requests
	clients = make(chan Client, size)
	// writing the results
	data = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

//  generate the required number of worker() goroutines for processing all requests
func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

// create all requests properly using the Client type and then to write them to the clients channel for processing
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// mimicking the client requests that you will have to process.
	go create(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%dtsquare: %d\n", d.job.integer, d.square)
		}
		// it does not allow the program to end until somebody writes something to the finished channel
		finished <- true
	}()

	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
