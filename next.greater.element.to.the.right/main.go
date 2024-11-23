package main

// tip: next greater element require monotonically decresing stack (i, e 9, 7, 1(where 1 being the top))
func nextLargerElement(arr []int, n int) []int {
	result := make([]int, n)

	for i := 0; i < len(arr); i++ {
		result[i] = -1
	}

	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, arr[i])

	}
	return result
}
