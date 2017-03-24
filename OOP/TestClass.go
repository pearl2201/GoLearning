package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Id   int
	Name string
}

func (person Person) print() string {
	return person.Name + " has id " + strconv.FormatInt(int64(person.Id), 10)
}

func main() {
	p := Person{Id: 13, Name: "Pearl"}
	fmt.Println(p.print())

}
