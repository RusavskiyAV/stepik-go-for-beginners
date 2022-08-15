package main

import "fmt"

const (
	degInHour = 30
	minInDeg  = 2
)

func main() {
	var d int

	if _, err := fmt.Scan(&d); err != nil {
		panic(err)
	}

	fmt.Print("It is", d/degInHour, "hours", d%degInHour*minInDeg, "minutes.")
}
