package main

import "fmt"

func main() {
	//双引号
	var name string = "simple"
	fmt.Printf("%T,%s\n", name, name)
	//	单引号 字符
	v1 := '1'
	v2 := 's'

	fmt.Printf("%T,%d\n", v1, v1)
	fmt.Printf("%T,%s\n", v2, v2)
	//	字符串拼接
	fmt.Println("hello " + "world")
}
