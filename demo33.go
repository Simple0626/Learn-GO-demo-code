package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i > 1 {
			goto exit
		}
	}
exit:
	fmt.Println("exec.")

}

func test() {

}
