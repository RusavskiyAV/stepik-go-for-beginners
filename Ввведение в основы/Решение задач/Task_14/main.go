package main

import (
	"fmt"
)

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Printf("%b", n)
}
