package main

import "fmt"

func main() {

	fmt.Println("2 + 3=", plus(2, 3))

	fmt.Println(multiValues())

	fmt.Println(variadicFunc(1, 2, 3, 4, 5))

	nums := []int{2, 3, 4, 5, 6}
	fmt.Println(variadicFunc(nums...))

	sli := make([]int, 3)
	copy(sli, nums)
	fmt.Println(variadicFunc(sli...))
	sli = append(sli, 1)
	fmt.Println(variadicFunc(sli...))
}

func plus(a int, b int) int {
	return a + b
}

func plusTwo(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multiValues() (int, int) {
	return 10, 200
}

func variadicFunc(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
