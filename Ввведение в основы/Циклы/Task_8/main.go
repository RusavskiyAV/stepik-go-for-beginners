package main

import (
	"fmt"
)

func main() {
	var (
		a, b   int
		digits []int
	)

	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	for tmp := a; tmp != 0; tmp /= 10 {
		digit := tmp % 10

		for tmp2 := b; tmp2 != 0; tmp2 /= 10 {
			if digit == tmp2%10 {
				digits = append(digits, digit)
			}
		}
	}

	l := len(digits)

	for i := l - 1; i >= 0; i-- {
		fmt.Print(digits[i])

		if i != 0 {
			fmt.Print(" ")
		}
	}
}
