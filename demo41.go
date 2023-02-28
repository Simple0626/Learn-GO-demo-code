package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 10: 9}
	s := data[:2:3]
	s = append(s, 23, 34)

	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
}
