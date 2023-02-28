package main

import "fmt"

func main() {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"sim1", 11},
		2: {"sim2", 12},
	}
	fmt.Println(m[1].name)
	fmt.Println(m[2].name)
}
