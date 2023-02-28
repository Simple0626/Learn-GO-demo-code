package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("传递前的数据：", arr)
	update1(arr)
	fmt.Println("传递后的数据：", arr)

}

func update1(arrs []int) {
	fmt.Println("传递的数据：", arrs)
	arrs[0] = 100
	fmt.Println("传递后的数据：", arrs)

}
