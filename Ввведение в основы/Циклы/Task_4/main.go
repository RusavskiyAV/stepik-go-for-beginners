package main

import (
	"fmt"
)

func main() {
	var n, count, max int

	for _, err := fmt.Scan(&n); n != 0; {
		if err != nil {
			panic(err)
		}

		if n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}

		_, err = fmt.Scan(&n)
	}

	fmt.Print(count)
}
