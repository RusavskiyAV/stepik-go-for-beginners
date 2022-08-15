package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var res int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		tmp, _ := strconv.Atoi(line)
		res += tmp
	}

	_, err := os.Stdout.WriteString(strconv.Itoa(res))

	if err != nil {
		panic(err)
	}
}
