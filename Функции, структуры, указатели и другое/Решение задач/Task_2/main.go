package main

import (
	"fmt"
)

func main() {
	var str string

	if _, err := fmt.Scan(&str); err != nil {
		panic(err)
	}

	r := []rune(str)

	for i, l := 0, len(r); i < l; i++ {
		fmt.Printf("%c", r[i])

		if i != l-1 {
			fmt.Print("*")
		}
	}
}
