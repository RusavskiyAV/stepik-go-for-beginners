package main

import (
	"fmt"
)

const minVal = 10

func main() {
	var n int

	for _, err := fmt.Scan(&n); n <= 100; {
		if err != nil {
			panic(err)
		}

		if n > minVal-1 {
			fmt.Println(n)
		}

		_, err = fmt.Scan(&n)
	}
}
