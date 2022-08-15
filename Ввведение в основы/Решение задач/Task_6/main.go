package main

import "fmt"

func main() {
	var a, b float32

	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	fmt.Print((a + b) / 2)
}
