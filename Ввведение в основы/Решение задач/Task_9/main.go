package main

import "fmt"

func getDigitalRoot(num int) int {
	var root int

	for ; num != 0; num /= 10 {
		root += num % 10
	}

	if root > 9 {
		return getDigitalRoot(root)
	} else {
		return root
	}
}

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	fmt.Print(getDigitalRoot(num))
}
