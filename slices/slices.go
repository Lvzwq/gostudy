package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("s:", s)

	fmt.Println("len(str):", len(s))
	fmt.Println("cap(str):", cap(s))

	s[0] = "a"
	s[1] = "bd"
	s[2] = "c"
	fmt.Println("s:", s)

	fmt.Println("len(str):", len(s))
	fmt.Println("cap(str):", cap(s))

	s = append(s, "d", "e")
	fmt.Println("s:", s)

	c := make([]string, 6)
	copy(c, s)
	fmt.Println("c:", c)

	// 包含头, 不包含尾部
	l := s[2:5]
	fmt.Println("l:", l)

	// 说明slice 尾部还有1个元素，为空
	l = s[5:]
	fmt.Println("l:", l)

	td := make([][]int, 3)
	fmt.Println("len(td):", len(td))
	for i := 0; i < len(td); i++ {
		inner := i + 1
		td[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			td[i][j] = i + j
		}
	}

	fmt.Println("td:", td)

	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7}
	sli := arr[1:4:5]
	fmt.Println(sli)
	fmt.Println("len(sli):", len(sli))
	fmt.Println("cap(sli):", cap(sli))
	sli = append(sli, 100)
	fmt.Println("len(sli):", len(sli))
	fmt.Println("cap(sli):", cap(sli))
	fmt.Println("arr:", arr)

	sli2 := arr[1:4:8]
	fmt.Println(sli2)

}
