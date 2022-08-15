package main

import (
	"fmt"
)

func main() {
	var a, b, res int

	for a >= b {
		if _, err := fmt.Scan(&a, &b); err != nil {
			panic(err)
		}
	}

	for i := a; i <= b; i++ {
		res += i
	}

	fmt.Print(res)
}
