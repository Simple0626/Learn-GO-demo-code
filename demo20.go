package main

import "fmt"

func main() {
	//	值传递
	//	定义数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	update(arr)
	fmt.Println(arr)
	//	不会有修改，这里的传递是一个copy

}

func update(arr [4]int) {
	arr[0] = 100
}
