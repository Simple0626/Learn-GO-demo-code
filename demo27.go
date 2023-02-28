package main

import "fmt"

func main() {

	sum := add(1, 2)
	fmt.Println(sum)
	sum2 := super(1, 2, add)
	fmt.Println(sum2)

	//调用匿名函数：
	sum3 := super(2, 3, func(a int, b int) int {
		return a + b
	})
	fmt.Println(sum3)
}

// 高阶函数：
func super(a, b int, f func(a int, b int) int) int {
	sum := f(a, b)
	return sum
}

func add(a, b int) int {
	return a + b
}
