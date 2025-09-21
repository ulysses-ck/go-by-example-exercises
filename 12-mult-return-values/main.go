package main

import (
	"fmt"
)

func vals() (int, int) {
	return 4, 2
}

func main() {
	bb, cc := vals()
	fmt.Println(bb)
	fmt.Println(cc)

	_, c := vals()
	fmt.Println(c)
}
