package main

import (
	"fmt"
)

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	fmt.Println(num % 10)
}
