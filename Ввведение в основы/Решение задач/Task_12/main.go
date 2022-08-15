package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i float64

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	maxPower := math.Log2(n)

	for i = 0; i <= maxPower; i++ {
		fmt.Print(math.Pow(2, i))

		if i != maxPower {
			fmt.Print(" ")
		}
	}
}
