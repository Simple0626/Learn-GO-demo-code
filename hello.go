package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	fmt.Println("hello world!")
	var x int     //自动初始化为0
	var y = false //自动推断为bool类型
	var name = "simple"
	name = "thanks"       //改变量
	var a, b int          //相同类型的多个变量
	var c, d = 100, "abc" //不同类型的初始化值,这里 c是100 d是abc
	/*
		按照编程习惯，建议以组的方式整理多行变量定义，即用大括号美观一点。
	*/
	var (
		a1, b1 int
		c1, d1 = 100, "abc"
	)
	fmt.Println(name, x, y, a, b, c, d)
	fmt.Println(a1, b1, c1, d1)

}
