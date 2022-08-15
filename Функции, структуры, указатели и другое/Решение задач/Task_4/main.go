package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var num string

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	for _, char := range num {
		tmp, _ := strconv.ParseFloat(string(char), 64)
		fmt.Print(math.Pow(tmp, 2))
	}
}
