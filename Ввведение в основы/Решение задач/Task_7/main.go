package main

import "fmt"

func main() {
	var n, num, count int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&num); err != nil {
			panic(err)
		}

		if num == 0 {
			count++
		}
	}

	fmt.Print(count)
}
