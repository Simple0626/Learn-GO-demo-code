package main

import "fmt"

func main() {
	var user = map[string]string{
		"name": "simple",
		"age":  "18",
		"sex":  "难",
	}
	if value, ok := user["father"]; ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("key not in map")
	}

}
