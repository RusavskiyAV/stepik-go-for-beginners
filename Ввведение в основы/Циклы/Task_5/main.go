package main

import (
	"fmt"
)

func main() {
	var n, c, d int

	if _, err := fmt.Scan(&n, &c, &d); err != nil {
		panic(err)
	}

	for i := 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)

			break
		}
	}
}
