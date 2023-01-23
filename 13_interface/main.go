package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	plate := Circle{3, 3, 10}
	mailbox := Rectangle{8, 10}

	fmt.Printf("Plate area: %f\n", getArea(plate))
	fmt.Printf("Mailbox area: %f\n", getArea(mailbox))
}
