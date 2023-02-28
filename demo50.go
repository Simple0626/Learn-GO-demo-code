package main

import "fmt"

type A struct {
	a int
	b int
}

type B struct {
	b float32
	c string
	d string
}

type C struct {
	A
	B
	a string
	c string
}

func main() {
	var c C
	c.a = "ca"
	fmt.Println(c.a)
	c.b = "ccc"
}
