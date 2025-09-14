package main

import "fmt"

func main() {
	s := []int{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	a := ""

	for i := range 11 {
		charRune := rune(s[i])
		charString := string(charRune)

		a += charString
	}

	fmt.Println(a)
}
