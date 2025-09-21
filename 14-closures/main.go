package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	asd := intSeq()

	fmt.Println(asd())
	fmt.Println(asd())
	fmt.Println(asd())

	newInts := intSeq()
	fmt.Println(newInts())
}
