package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		if _, err := fmt.Scan(&workArray[i]); err != nil {
			panic(err)
		}
	}

	var a, b int

	for i := 0; i < 3; i++ {
		if _, err := fmt.Scan(&a, &b); err != nil {
			panic(err)
		}

		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}
