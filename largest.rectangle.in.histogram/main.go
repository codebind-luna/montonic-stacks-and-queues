package main

import "math"

func nextSmallerToTheLeft(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	for i := 0; i < len(arr); i++ {
		result[i] = -1
	}

	stack := [][]int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[i] <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1][1]
		}

		stack = append(stack, []int{arr[i], i})

	}
	return result
}

func nextSmallerToTheSRight(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := 0; i < len(arr); i++ {
		result[i] = -1
	}

	stack := [][]int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1][1]
		}

		stack = append(stack, []int{arr[i], i})

	}
	return result
}

func largestRectangleArea(heights []int) int {
	left := nextSmallerToTheLeft(heights)
	right := nextSmallerToTheSRight(heights)

	maxArea := math.MinInt

	for i := 0; i < len(heights); i++ {
		maxArea = max(maxArea, (right[i]-left[i]-1)*heights[i])
	}

	return maxArea
}
