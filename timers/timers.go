package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second)
	fmt.Println(time.Now())
	nowTime := <-time1.C
	fmt.Println(nowTime)

	p := fmt.Println
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	now := time.Now()
	fmt.Println("now", now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Weekday())

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())

	p(now.Unix())
	p(now.UnixMilli())

	var a int32 = 12
	var b int32 = 34
	c := a + b
	fmt.Printf("%t", c)
}
