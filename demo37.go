package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := data[1:4:8]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
