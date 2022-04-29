package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("m:", m)

	v1, c1 := m["k1"]
	fmt.Println("key:", v1, ", c:", c1)

	delete(m, "k1")
	v2, c2 := m["k1"]
	fmt.Println("key:", v2, ", c:", c2)

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for key := range kvs {
		fmt.Println("key:", key)
	}

	for key, value := range kvs {
		fmt.Println("key:", key, ", value:", value)
	}

}
