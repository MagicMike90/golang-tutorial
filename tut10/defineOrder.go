package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

//  The function is blocked until the a channel is closed
func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

// The C() function is blocked and waits for the a channel to close in order to start working.
func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}
func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}
