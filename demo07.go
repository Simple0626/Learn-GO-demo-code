package main

import "fmt"

func main() {
	const (
		a = iota
		b = 3
		c = iota
		d
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)

	const (
		g = iota
	)
	fmt.Println(g)
}
