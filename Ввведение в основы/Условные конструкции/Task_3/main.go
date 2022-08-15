package main

import "fmt"

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	digit := -1

	for tmp := num; tmp != 0; tmp /= 10 {
		digit = tmp % 10
	}

	fmt.Print(digit)
}
