package main

import "fmt"

func main() {
	f("1")
	fmt.Println("2")
	defer f("3")
	fmt.Println("4")
	defer f("5")
	fmt.Println("6")
}
func f(s string) {
	fmt.Println(s)
}
