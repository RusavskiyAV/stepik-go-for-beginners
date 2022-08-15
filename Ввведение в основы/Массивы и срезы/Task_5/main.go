package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	var tmp, count int

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&tmp); err != nil {
			panic(err)
		}

		if tmp > 0 {
			count++
		}
	}

	fmt.Print(count)
}
