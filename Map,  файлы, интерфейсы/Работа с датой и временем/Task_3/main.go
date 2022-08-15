package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const format = "2006-01-02 15:04:05"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	dates := strings.Split(scanner.Text(), ",")
	date1, _ := time.Parse(format, dates[0])
	date2, _ := time.Parse(format, dates[1])

	if date1.Before(date2) {
		fmt.Print(date2.Sub(date1))
	} else {
		fmt.Print(date1.Sub(date2))
	}
}
