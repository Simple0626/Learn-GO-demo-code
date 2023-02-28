package main

import "fmt"

func main() {
	const (
		//在常量组中如不能指定类型和初始化值，则与上一行非空常量右值（表达式文本）相同。
		x uint16 = 120
		y        //与上一行x类型，右值相同
		s = "abc"
		z //与s的类型，右值相同
	)

	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
}
