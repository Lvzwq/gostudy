package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan uint)

	go func() {
		var i uint
		for i = 0; i < 10; i++ {
			message <- i
		}
	}()

	go func() {
		for {
			if value, ok := <-message; ok {
				fmt.Print(ok, value, "\t")
			} else {
				fmt.Println(ok, "empty result")
			}
		}
	}()

	time.Sleep(time.Second * 1)

	fmt.Println()

	messageBuffer := make(chan string, 2)
	messageBuffer <- "hello"
	messageBuffer <- "world"

	fmt.Println(<-messageBuffer)
	fmt.Println(<-messageBuffer)
	messageBuffer <- "go"
	fmt.Println(<-messageBuffer)

	//fmt.Println(<-messageBuffer)

}
