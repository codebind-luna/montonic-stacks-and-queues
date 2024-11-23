package main

func prevSmaller(A []int) []int {
	n := len(A)
	result := make([]int, n)

	for i := 0; i < len(A); i++ {
		result[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= A[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, A[i])

	}
	return result
}
