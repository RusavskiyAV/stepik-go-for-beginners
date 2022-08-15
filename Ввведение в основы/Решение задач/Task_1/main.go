package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Print(n / 100 + n / 10 % 10 + n % 10)
}
