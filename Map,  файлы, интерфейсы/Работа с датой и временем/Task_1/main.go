package main

import (
	"fmt"
	"time"
)

func main() {
	var str string

	if _, err := fmt.Scan(&str); err != nil {
		panic(err)
	}

	now, _ := time.Parse(time.RFC3339, str)
	fmt.Print(now.Format(time.UnixDate))
}
