package main

import "fmt"

func main() {
	a := add()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	var fib func(int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
}

func add() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
