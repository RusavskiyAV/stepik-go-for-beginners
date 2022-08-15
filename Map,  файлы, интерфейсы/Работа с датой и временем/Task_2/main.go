package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	dinner = 13
	format = "2006-01-02 15:04:05"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	now, _ := time.Parse(format, scanner.Text())

	if now.Hour() < dinner {
		fmt.Print(now.Format(format))
	} else {
		fmt.Print(now.Add(time.Hour * 24).Format(format))
	}
}
