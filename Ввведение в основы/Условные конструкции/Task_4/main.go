package main

import (
	"fmt"
	"unicode"
)

const strLen = 6

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func sum(s string) int {
	l := len(s)
	var sum int

	for i := 0; i < l; i++ {
		sum += int(s[i])
	}

	return sum
}

func main() {
	var num string

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	if len(num) != strLen || !isInt(num) {
		panic("некорректное число")
	}

	if sum(num[:strLen/2]) == sum(num[strLen/2:]) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
