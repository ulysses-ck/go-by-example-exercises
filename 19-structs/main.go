package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	fmt.Println(person{"falo", 20})

	fmt.Println(person{name: "aldo", age: 30})

	fmt.Println(person{name: "alberto"})

	fmt.Println(newPerson("kici"))

	s := person{name: "maxi", age: 7}

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"axel",
		true,
	}
	fmt.Println(dog)
}
