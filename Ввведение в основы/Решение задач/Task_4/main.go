package main

import "fmt"

func main() {
	var a, b, c int

	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		panic(err)
	}

	if a*a+b*b == c*c {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
