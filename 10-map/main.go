package main

import (
	"fmt"
	"maps"
)

func main() {
	c := make(map[string]int)

	c["thekey"] = 42
	c["troll"] = 13

	fmt.Println("map:", c)

	v1 := c["thekey"]
	fmt.Println("v1: ", v1)

	cc := make(map[string]int)
	theKey := "theKey"
	troll := "troll"

	cc[theKey] = 42
	cc[troll] = 13 + 6

	fmt.Println("map:", cc)

	delete(cc, troll)
	fmt.Println("map:", cc)

	clear(cc)
	fmt.Println("map:", cc)

	_, prs := cc[troll]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 32678}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 32678}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
