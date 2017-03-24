package main

import (
	"fmt"
)

type Matrix [][]bool

const (
	INFINITY = 100
	NO_TRACE = -1
)

func main() {
	board := Matrix{
		{false, false, true, false},
		{false, false, true, true},
		{true, true, false, false},
		{false, true, false, false}}
	start := 0
	end := 3
	size := 4
	result := Init(board, start, end, size)
	if result == nil {
		fmt.Println("can not found a line")
	} else {
		for _, v := range result {
			fmt.Printf("%d ", v)
		}
	}

}

func Init(board Matrix, start, end, size int) []int {
	var result []int
	var stack []int
	var point int
	trace := make([]int, size)
	for i := 0; i < size; i++ {
		trace[i] = NO_TRACE
	}
	trace[start] = INFINITY
	stack = append(stack, start)
	for len(stack) > 0 {
		point = stack[len(stack)-1]
		isFound := false
		for i := 0; i < size && !isFound; i++ {
			if board[point][i] && trace[i] == NO_TRACE {
				trace[i] = point
				stack = append(stack, i)
				isFound = true

			}
		}
		if !isFound {
			stack = stack[:len(stack)-1]
		}
	}

	if trace[end] == -1 {
		return nil
	} else {
		last := end
		for last != start {
			result = append(result, last)
			last = trace[last]
		}
		result = append(result, last)
	}
	return result
}
