package main

import "fmt"

func main() {
	var name, email = "Foo", "foo@example.com"
	var age = 31
	var height float32 = 172.3

	fmt.Println("Name:", name, "\nEmail:", email, "\nAge:", age, "\nHeight:", height)
	fmt.Printf("%T\n", age)
}
