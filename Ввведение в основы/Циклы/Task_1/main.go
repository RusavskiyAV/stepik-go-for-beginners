package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64

	for i = 1; i <= 10; i++ {
		fmt.Println(math.Pow(i, 2))
	}
}
