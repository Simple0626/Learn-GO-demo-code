package main

import "fmt"

func add(a, b int) int {
	c := a + b
	return c
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
}
