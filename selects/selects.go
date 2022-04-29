package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan int)

	go func() {
		c1 <- "Hello C1"
	}()

	go func() {
		c2 <- 2022
	}()

	m := make(map[string]interface{})

	for i := 0; i < 2; i++ {
		select {
		case c1Resp := <-c1:
			m["c1"] = c1Resp
			fmt.Println(c1Resp)
		case c2Resp := <-c2:
			fmt.Println(c2Resp)
			m["c2"] = c2Resp
		}
	}
	fmt.Println(m)

	a1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		a1 <- "result 1"
	}()

	select {
	case m1 := <-a1:
		fmt.Println(m1)
	case <-time.After(time.Second):
		fmt.Println("a1 timeout")
	}

	time.Sleep(3 * time.Second)

	msg := "hi"
	messages := make(chan string)
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	case mess := <-messages:
		fmt.Println("message received", mess)
	default:
		fmt.Println("no message sent")
	}
}
