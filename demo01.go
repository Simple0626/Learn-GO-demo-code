package main

import "fmt"

//func main() {
//	x := 1
//	name := "simple"
//	age := 18
//	fmt.Println(x)
//	fmt.Println(name)
//	fmt.Println(age)
//}

/*func main() {
	x := 100
	println(&x)

	x := 200 //错误：no new variable on left side of :=
	println(&x, x)
}*/

func main() {
	x := 100
	//输出变量的内存地址和值
	println(&x, x)
	fmt.Printf("%d,%p", x, &x)
	println()
	//得出变量的类型
	fmt.Printf("%T", x)
}
