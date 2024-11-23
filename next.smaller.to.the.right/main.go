package main

import "fmt"

func nextSmallerToRight(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := 0; i < len(arr); i++ {
		result[i] = -1
	}

	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, arr[i])

	}
	return result
}

func main() {
	arr := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(arr)
	fmt.Println(nextSmallerToRight(arr))
}
