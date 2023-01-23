package main

import "fmt"

func main() {
	countryCode := map[string]string{"US": "United States", "UK": "United Kingdom", "ID": "Indonesia"}

	delete(countryCode, "US")

	countryCode["DE"] = "Germany"

	fmt.Println(countryCode)
}
