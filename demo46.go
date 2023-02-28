package main

import "fmt"

func main() {
	var user = map[string]string{
		"name": "simple",
		"age":  "18",
		"sex":  "难",
	}
	//增加
	user["father"] = "complex"

	fmt.Println(user)
	//	删除
	delete(user, "father")
	fmt.Println(user)
}
