package main

import "fmt"

func main() {
	x := 1234
	p := &x
	//p++ // Error: invalid operation: p += 1 (mismatched types *int and int)
	fmt.Println(p)
}
