package main

import (
	"fmt"
)

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	if 0 == n {
		fmt.Print(0)

		return
	}

	prev, next, i := 0, 1, 1

	for next <= n {
		if next == n {
			fmt.Print(i)

			return
		}

		prev, next = next, prev+next
		i++
	}

	fmt.Print(-1)
}
