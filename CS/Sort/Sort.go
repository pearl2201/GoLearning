package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var n int = 10

	// buble sort

	BubbleSort(n)

	SelectionSort(n)

	MergerSort(n)

	QuickSort(n)
}

func BubbleSort(n int) {
	fmt.Println("----------Bubble Sort--------------")
	var arr = makeArr(n)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}

	}
	printArr(arr)

}

func SelectionSort(n int) {
	fmt.Println("----------Selection Sort--------------")
	var arr = makeArr(n)
	for i := 0; i < n-1; i++ {
		indexMin := n - 1
		for j := n - 2; j >= i; j-- {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
		}
		if indexMin != i {
			temp := arr[i]
			arr[i] = arr[indexMin]
			arr[indexMin] = temp
		}
	}
	printArr(arr)

}

func MergerSort(n int) {
	fmt.Println("----------Merger Sort--------------")
	var arr = makeArr(n)
	arr = MergerSlice(arr)
	printArr(arr)
}

func MergerSlice(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	left = MergerSlice(left)
	right = MergerSlice(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])

			j++
		}

	}

	if i < len(left) {
		for ; i < len(left); i++ {
			result = append(result, left[i])
		}
	}

	if j < len(right) {
		for ; j < len(right); j++ {
			result = append(result, right[j])
		}
	}
	return result
}

func QuickSort(n int) {
	fmt.Println("----------Quick Sort--------------")
	var arr = makeArr(n)
	arr = quickSlice(arr)
	printArr(arr)
}
func quickSlice(m []int) []int {

	if len(m) <= 1 {
		return m
	}
	key := m[len(m)/2]
	var left, right, middle []int
	for _, v := range m {
		if v < key {
			left = append(left, v)
		} else if v == key {
			middle = append(middle, v)
		} else {
			right = append(right, v)
		}

	}
	left = quickSlice(left)
	right = quickSlice(right)

	left = append(left, middle...)
	left = append(left, right...)
	return left
}

func makeArr(size int) []int {

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(90)
	}
	printArr(arr)

	fmt.Println("----")
	return arr

}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n")
}
