package main

import "fmt"

func main() {
	var fruits [3]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"

	sizes := [3]string{"Small", "Medium", "Average"}
	foodSlice := []string{"Pizza", "Burger", "Chicken", "Pop Corn"}

	fmt.Println(fruits)
	fmt.Println(fruits[2])
	fmt.Println(sizes)
	fmt.Println(foodSlice)
	fmt.Println(len(foodSlice))
	fmt.Println(foodSlice[1:3])
}
