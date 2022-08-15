package main

import "fmt"

func main() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		panic(err)
	}

	fmt.Print("It is ", a/3600, " hours ", a%3600/60, " minutes.")
}
