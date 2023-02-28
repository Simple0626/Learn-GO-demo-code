package main

import "fmt"

func main() {
	j := 5
	a := func() func() {
		i := 1
		return func() {
			fmt.Println(j, i)
		}
	}()
	a()
	j *= 2
	a()
	
}
