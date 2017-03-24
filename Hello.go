package main

import "fmt"

func main() {
	fmt.Printf("Hello, world\n")

	var x, y, z = 3, 4, "foo"
	fmt.Printf("x is of the type %T\n", x)
	fmt.Printf("y is of the type %T\n", y)
	fmt.Printf("z is of the type %T\n", z)

	var a int = 10

LOOP:
	{
		for a < 20 {
			if a == 15 {
				a = a + 1
				goto LOOP
			}
			fmt.Printf("Value of a: %d\n", a)
			a++
		}
	}
	fmt.Printf("Value of max(3,4): %d\n", max(3, 4))
	var k = [6]int{4, 2, 1, 5, 6, 3}
	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			if k[j] < k[j-1] {
				k[j], k[j-1] = swap(k[j], k[j-1])
			}
		}
	}
	for i := 0; i < 6; i++ {
		fmt.Printf("%d\n", k[i])
	}

	fmt.Printf("Value of pos(0): %d\n", *(&k[0]))
}

func swap(x, y int) (int, int) {
	return y, x
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
