package main

import (
	"fmt"
)

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	if num%400 == 0 || num%4 == 0 && num%100 != 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
