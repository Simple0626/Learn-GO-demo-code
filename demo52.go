package main

import "fmt"

type City string

type (
	A1 int
	A2 int
)

func main() {
	myCity := City("beijing")
	fmt.Println(myCity)

	//fmt.Println("A1 is ", A1(1))
	//fmt.Println("A2 is ", A2(2))

	//if !(A2(2) >= A1(1)) {
	//	fmt.Println("yes")
	//}

}
