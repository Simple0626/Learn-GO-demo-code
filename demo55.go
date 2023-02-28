package main

import (
	"errors"
	"fmt"
)

func panicFunc() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover panic")
		}
	}()
	panic(errors.New("test is for panic"))
}

func main() {
	fmt.Println("before panic")
	panicFunc()
	fmt.Println("after panic")
}
