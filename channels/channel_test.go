package main

import (
	"fmt"
	"testing"
	"time"
)

func Worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func TestChannelSync(t *testing.T) {
	done := make(chan bool)
	go Worker(done)
	<-done
}

func TestSync(t *testing.T) {
	fmt.Println("hello testing")
}

func TestCurrentSync(t *testing.T) {
	c1 := make(chan string, 2)
	c1 <- "hello"
	c1 <- "world"

	close(c1)
	for elem := range c1 {
		fmt.Println(elem)
	}
}

func TestCurAsync(t *testing.T) {
	c1 := make(chan string, 2)
	go func() {
		c1 <- "hello"
		c1 <- "world"
	}()

	for elem := range c1 {
		fmt.Println(elem)
	}

	close(c1)
}
