package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) sayhello() {
	fmt.Println("hello student and method")

}

//func (s Student) name() {
//
//}

func (s *Student) addAge() {
	s.age += 1
}
func (s Student) printAge() {
	fmt.Println(s.name, "的年龄是：", s.age)
}
func main() {
	s := Student{"simple", 18}
	s.sayhello()
	s.addAge()
	s.printAge()

}
