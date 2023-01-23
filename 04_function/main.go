package main

import "fmt"

func main() {
	fmt.Println(greetings("Foo"))
	fmt.Println(sumOfTwo(3, 1))
}

func greetings(name string) string {
	return "Welcome back " + name
}

func sumOfTwo(num1, num2 int) int {
	return num1 + num2
}
