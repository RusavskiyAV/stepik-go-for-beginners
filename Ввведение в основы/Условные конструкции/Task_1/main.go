package main

import "fmt"

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	switch {
	case num > 0:
		fmt.Print("Число положительное")
	case num < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
