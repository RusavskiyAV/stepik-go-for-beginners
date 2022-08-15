package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	var tmp int

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&tmp); err != nil {
			panic(err)
		}

		if i%2 == 0 {
			fmt.Print(tmp)

			if i < n-1 {
				fmt.Print(" ")
			}
		}
	}
}
