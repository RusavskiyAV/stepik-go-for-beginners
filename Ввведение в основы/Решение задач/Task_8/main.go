package main

import "fmt"

func main() {
	var n, num, count, min int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&num); err != nil {
			panic(err)
		}

		if i == 0 {
			min = num
		}

		if num < min {
			min = num
			count = 1
		} else if num == min {
			count++
		}
	}

	fmt.Print(count)
}
