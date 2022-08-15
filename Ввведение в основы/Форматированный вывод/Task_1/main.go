package main

import (
	"fmt"
	"math"
)

func main() {
	var a, r float64

	if _, err := fmt.Scan(&a); err != nil {
		panic(err)
	}

	r = math.Pow(a, 2)

	switch {
	case a <= 0:
		fmt.Printf("число %2.2f не подходит", a)
	case r > 10000:
		fmt.Printf("%e", a)
	default:
		fmt.Printf("%.4f", r)
	}
}
