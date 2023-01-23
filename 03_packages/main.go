package main

import (
	"fmt"
	"math"

	"github.com/mirhasalh/go_example/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(172.3))
	fmt.Println(math.Ceil(172.3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Hello"))
}
