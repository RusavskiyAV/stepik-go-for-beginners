package main

import (
	"fmt"
)

func main() {
	var str string

	if _, err := fmt.Scan(&str); err != nil {
		panic(err)
	}

	max := str[0]

	for i, l := 0, len(str); i < l; i++ {
		digit := str[i]

		if max < digit {
			max = digit
		}
	}

	fmt.Print(string(max))
}
