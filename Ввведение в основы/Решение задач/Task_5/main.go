package main

import "fmt"

func main() {
	var a, b, c int

	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		panic(err)
	}

	if a+b > c && b+c > a && a+c > b {
		fmt.Print("Существует")
	} else {
		fmt.Print("Не существует")
	}
}
