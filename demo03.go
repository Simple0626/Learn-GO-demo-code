package main

import "fmt"

func test() (int, int) {
	return 100, 200
}
func main() {
	a, _ := test()
	fmt.Println(a)

}
