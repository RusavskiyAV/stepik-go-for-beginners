package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	fmt.Print(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}
