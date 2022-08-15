package main

import (
	"fmt"
)

func main() {
	var (
		i          int
		t, x, p, y float32
	)

	if _, err := fmt.Scan(&x, &p, &y); err != nil {
		panic(err)
	}

	for t = x; t < y; i++ {
		t += t / 100.0 * p
	}

	fmt.Print(i)
}
