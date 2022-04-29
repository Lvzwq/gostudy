package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i, time.Now())
		time.Sleep(time.Duration(i) * time.Second)
	}
}

func main() {
	go f("goroutine")

	f("direct")

	go func(msg string) {
		fmt.Println("go func, ", msg)
	}("goroutine")
	fmt.Println("go func direct")
	time.Sleep(time.Second)

}
