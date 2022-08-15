package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type average struct {
	Average float32
}

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type myStruct struct {
	Id       int
	Number   string
	Year     int
	Students []student
}

func main() {
	var s myStruct

	data, _ := ioutil.ReadAll(os.Stdin)

	if err := json.Unmarshal(data, &s); err != nil {
		panic(err)
	}

	var c int

	for _, w := range s.Students {
		c += len(w.Rating)
	}

	res, _ := json.MarshalIndent(average{float32(c) / float32(len(s.Students))}, "", "    ")

	fmt.Print(string(res))
}
