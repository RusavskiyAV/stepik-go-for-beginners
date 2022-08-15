package main

import (
	"fmt"
	"math"
)

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	fmt.Print(math.Pow(float64(num), 2))
}
