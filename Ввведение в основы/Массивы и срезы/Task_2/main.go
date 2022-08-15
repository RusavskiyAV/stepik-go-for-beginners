package main

import "fmt"

const index = 3

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			panic(err)
		}
	}

	fmt.Print(arr[index])
}
