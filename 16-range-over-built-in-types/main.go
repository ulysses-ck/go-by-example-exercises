package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"l": "linux", "p": "potato"}
	for k, v := range kvs {
		s1 := fmt.Sprintf("%s -> %s", k, v)
		fmt.Println(s1)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
