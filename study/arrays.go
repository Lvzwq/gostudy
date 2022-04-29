package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("a:", a)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)
	fmt.Println("len(b):", len(b))

	c := [][]int{{1, 2}, {3, 4}}
	fmt.Println("c:", c)
}
