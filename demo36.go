package main

import "fmt"

func main() {
	a := 100
	b := 200
	exchange(&a, &b)
	fmt.Println(a, b)
}

func exchange(a, b *int) (int, int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
	return *a, *b

}
