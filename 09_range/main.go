package main

import "fmt"

func main() {
	ids := []int{1, 5, 4}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// SUM up the ids
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	countryCode := map[string]string{"US": "United States", "UK": "United Kingdom", "ID": "Indonesia"}

	for k, v := range countryCode {
		fmt.Printf("%s: %s\n", k, v)
	}
}
