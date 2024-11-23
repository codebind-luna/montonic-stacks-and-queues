package main

import "fmt"

func nextGreater(height []int) []int {
	n := len(height)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && height[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}

		stack = append(stack, height[i])
	}

	return ans
}

func nearestGreaterToLeft(height []int) []int {
	n := len(height)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}

		stack = append(stack, height[i])
	}

	return ans
}

func trap(height []int) int {
	n := len(height)
	right := nextGreater(height)
	left := nearestGreaterToLeft(height)

	maxArea := 0

	for i := 0; i < n; i++ {
		rH := 0
		lH := 0

		if right[i] != -1 {
			rH = right[i]
		}

		if left[i] != -1 {
			lH = left[i]
		}

		minH := min(lH, rH)

		if minH > height[i] {
			maxArea = max(maxArea, minH-height[i])
		}
	}

	return maxArea
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(height)
	trap(height)
}
