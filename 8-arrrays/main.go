package main

import "fmt"

func main() {
	var a [6]int

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [...]int{2, 3, 4, 5, 6}
	fmt.Println("dcl", c)

	d := [...]int{100, 3: 400, 500}
	fmt.Println("idx", d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
