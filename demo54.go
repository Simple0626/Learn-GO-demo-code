package main

import "fmt"

func t() {
	defer fmt.Println("world")
	fmt.Print("hello ")
}
func main() {
	fmt.Println("simple")
	t()
}
