package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("sum:", sum)

	for index, value := range nums {
		if value == 3 {
			fmt.Println("index=", index)
		}
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
