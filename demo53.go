package main

import (
	"fmt"
	"math"
)

// 定义接口
type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	len, width float32
}

// 实现接口中的方法
func (r *rect) area() float32 {
	return r.len * r.width
}

func (r *rect) perim() float32 {
	return 2 * (r.len + r.width)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	//TODO implement me
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float32 {
	//TODO implement me
	return 2 * math.Pi * c.radius
}

func show(name string, params interface{}) {
	switch params.(type) {
	case geometry:
		fmt.Printf("area of %v is %v \n", name, params.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, params.(geometry).perim())
	default:
		fmt.Println("wrong type")
	}

}

func main() {
	r := &rect{
		len:   1,
		width: 2,
	}
	show("rect", r)
	c := &circle{
		radius: 1,
	}
	show("cicle", c)

}
