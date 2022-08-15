package main

import "fmt"

func main() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		panic(err)
	}

	fmt.Print(a/100 + a/10%10*10 + a%10*100)
}
