package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{}
	p2 := &person{}
	p3 := new(person)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
