package main

import "fmt"

func main() {
	f := add
	fmt.Println(f()(1))
	fm := add()
	fmt.Println(fm(10))
	fmt.Println(fm(100))
	fmt.Println(fm(1000))
}

func add() func(i int) int {
	x := 0
	return func(i int) int {
		x += i
		return x
	}

}
