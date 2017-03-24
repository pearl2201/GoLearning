package main

import (
	"fmt"
)

func t(x int) func() int {

	return func() int {

		return x
	}
}

func main() {

	var whatever [5]struct{}

	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}

	for i := range whatever {
		defer func(n int) { fmt.Println(n) }(i)
	}
}
