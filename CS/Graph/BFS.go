package main

import (
	"fmt"
)

type Matrix [][]bool

const (
	NO_TRACE    = -1
	START_TRACE = -2
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
	result := BFS(board, start, end, size)
	if result == nil {
		fmt.Println("can not found a line")
	} else {
		for _, v := range result {
			fmt.Printf("%d ", v)
		}
	}

}

func BFS(board Matrix, start, end, size int) []int {
	var result []int

	trace := make([]int, size)
	for i := 0; i < size; i++ {
		trace[i] = NO_TRACE
	}
	trace[start] = START_TRACE
	var currArr, oldArr []int
	currArr = append(currArr, start)
	for len(currArr) > 0 {
		for i := 0; i < len(currArr); i++ {
			for j := 0; j < size; j++ {
				if board[currArr[i]][j] && trace[j] == NO_TRACE {
					trace[j] = currArr[i]
					oldArr = append(oldArr, j)
				}
			}
		}
		if len(oldArr) == 0 {
			break

		} else {
			currArr = oldArr
			oldArr = nil
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
