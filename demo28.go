package main

import "fmt"

func main() {
	f2 := increment()
	fmt.Println(f2())
}

func increment() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}
