package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Float64())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
}
