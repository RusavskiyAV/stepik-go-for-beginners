package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	duration, _ := time.ParseDuration(strings.ReplaceAll(strings.Replace(strings.Replace(scanner.Text(), "мин.", "m", 1), "сек.", "s", 1), " ", ""))
	fmt.Print(time.Unix(now, 0).UTC().Add(duration).Format(time.UnixDate))
}
