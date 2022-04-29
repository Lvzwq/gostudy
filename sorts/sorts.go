package main

import (
	"fmt"
	"sort"
)

type CustomSort []string

func (customSort CustomSort) Less(i, j int) bool {
	return len(customSort[i]) < len(customSort[j])
}

func (customSort CustomSort) Swap(i, j int) {
	customSort[i], customSort[j] = customSort[j], customSort[i]
}

func (customSort CustomSort) Len() int {
	return len(customSort)
}

func main() {
	ints := []int{1, 10, 121, 12}
	sort.Ints(ints)
	fmt.Println("ints", ints)

	strs := []string{"ba", "acde", "abc"}
	fmt.Println(sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println("strs", strs)
	fmt.Println(sort.StringsAreSorted(strs))

	strc := []string{"ba", "acde", "abc"}

	sort.Sort(CustomSort(strc))
	fmt.Println(strc)
}
