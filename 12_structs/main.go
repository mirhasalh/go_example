package main

import (
	"fmt"
	"strconv"
)

type Fruit struct {
	name   string
	weight float32
	qty    int
}

func (f Fruit) sayHi() string {
	return "Hi, I'm " + f.name + " and my qty is " + strconv.Itoa(f.qty)
}

func main() {
	banana := Fruit{
		name:   "Banana",
		weight: 2.1,
		qty:    12,
	}

	apple := Fruit{"Apple", 1.8, 10}

	fmt.Println(banana, apple)
	fmt.Println(banana.name)
	fmt.Println(apple.sayHi())
}
