package main

import "fmt"

func main() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(a)
	fmt.Println(b)
}
