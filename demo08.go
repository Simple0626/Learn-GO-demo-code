package main

import "fmt"

func main() {
	var age int = 18

	fmt.Println(age)
	fmt.Printf("%T", age)

	var money float64 = 3.14
	fmt.Printf("%T,%.1f", money, money)

}
