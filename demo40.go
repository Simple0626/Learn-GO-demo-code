package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)

	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s, s2)

	s3 := append(s2, 2, 3) //可以添加多个值
	fmt.Println(s3)
}
