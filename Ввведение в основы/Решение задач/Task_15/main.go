package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num, r int

	if _, err := fmt.Scan(&num, &r); err != nil {
		panic(err)
	}

	fmt.Printf(strings.ReplaceAll(strconv.Itoa(num), strconv.Itoa(r), ""))
}
