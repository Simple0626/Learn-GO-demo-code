package main

import (
	"fmt"
)

func main() {
	sum := getSum1(5)
	fmt.Println(sum)
}
func getSum1(n int) (sum int) {
	if n == 1 {
		return 1
	}
	return getSum1(n-1) + n
}
