package main

import "fmt"

func main() {
	//	f1()加了括号是函数调用
	fmt.Printf("%T\n", f1)
	f2 := f1
	fmt.Printf("%T\n", f2)

	//var f3 func(int, int) (int, int)
	//f3 = f1
	//fmt.Printf("%T\n", f3)

	var f4 func(int, int) = f1
	f4(1, 2)

}

func f1(a, b int) {
	fmt.Println(a, b)
}
