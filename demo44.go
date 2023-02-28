package main

import "fmt"

func main() {
	var a = map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(a[1])
	var b = map[string]string{}
	c := map[string]bool{}
	d := make(map[string]int)
	var e = map[string]string{
		"name": "simple",
		"age":  "18",
		"sex":  "难",
	}
	fmt.Println(e["name"], e["age"])
	fmt.Println(e["jsjsjgsl"]) //输出空字符串
}
