package main

import (
	"fmt"
)

func main() {
	var n, num, sum int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&num); err != nil {
			panic(err)
		}

		if num/10 > 0 && num/100 == 0 && num%8 == 0 {
			sum += num
		}
	}

	fmt.Print(sum)
}
