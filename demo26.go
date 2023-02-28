package main

import "fmt"

func main() {
	f1()
	f2 := f1
	f2()

	//	匿名函数
	f3 := func() {
		fmt.Println("f3函数")
	}
	f3()

	func() {
		fmt.Println("f4函数")
	}()

	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("f5函数")
	}(1, 2)

	sum := func(a, b int) int {
		fmt.Println("f6函数")
		return a + b

	}(1, 2)
	fmt.Println(sum)

}

func f1() {
	fmt.Println("f1函数")
}
