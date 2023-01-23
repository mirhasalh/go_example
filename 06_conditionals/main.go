package main

import "fmt"

func main() {
	x := 13
	y := 6

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	brightness := "light"

	switch brightness {
	case "light":
		fmt.Println("Brightness is light")
	case "dark":
		fmt.Println("Brightness is dark")
	default:
		fmt.Println("Undifined")
	}
}
