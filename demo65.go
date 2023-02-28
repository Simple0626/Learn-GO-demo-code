package main

import "encoding/json"

type House struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	house := House{
		Name: "sim",
		Age:  100,
	}
	jsonhouse, err := json.Marshal(&house)
	if err != nil {
		println("no")
	}
	println(string(jsonhouse))
}
