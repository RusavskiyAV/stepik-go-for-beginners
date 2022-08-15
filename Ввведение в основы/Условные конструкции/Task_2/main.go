package main

import "fmt"

func main() {
	var num int

	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}

	for tmp := num; tmp != 0; tmp /= 10 {
		digit := tmp % 10

		for tmp2 := tmp / 10; tmp2 != 0; tmp2 /= 10 {
			if digit == tmp2%10 {
				fmt.Print("NO")

				return
			}
		}
	}

	fmt.Print("YES")
}
