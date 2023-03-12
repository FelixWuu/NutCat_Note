package main

import (
	"fmt"
	"math/rand"
	"time"
)

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func add(c chan int) {
	sum, count := 0, 0
	t := time.After(1 * time.Second)

	for {
		select {
		case input := <-c:
			sum += input
			count += 1

		case <-t:
			c = nil
			fmt.Printf("sum %d, count %d", sum, count)
		}
	}
}

func NilChannelDemo() {
	c := make(chan int)
	go add(c)
	go send(c)
	time.Sleep(3 * time.Second)
}
