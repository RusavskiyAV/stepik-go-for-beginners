package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	fmt.Print(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
}
