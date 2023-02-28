package main

import "fmt"

func main() {
	name := "simple"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
}
