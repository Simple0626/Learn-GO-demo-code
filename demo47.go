package main

import "fmt"

func main() {
	var user = map[string]string{
		"name": "simple",
		"age":  "18",
		"sex":  "难",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}
}
