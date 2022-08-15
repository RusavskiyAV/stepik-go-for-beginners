package main

import "fmt"

const delimiter = 7

func main() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil && a > b {
		panic(err)
	}

	for ; b >= a; b-- {
		if b%delimiter == 0 {
			fmt.Print(b)

			return
		}
	}

	fmt.Print("NO")
}
