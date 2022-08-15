package main

import "fmt"

const count = 5

func main() {
	var n int

	array := [count]int{}

	for i := 0; i < count; i++ {
		if _, err := fmt.Scan(&n); err != nil {
			panic(err)
		}

		array[i] = n
	}

	max := array[0]

	for i := 1; i < count; i++ {

		if max < array[i] {
			max = array[i]
		}
	}

	fmt.Print(max)
}
