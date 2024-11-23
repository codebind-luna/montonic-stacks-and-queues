/* Leetcode Problem: https://leetcode.com/problems/next-greater-element-i/ */

package main

func nextLargerElement(arr []int, n int, m map[int]int, result []int) {
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		} else {
			result[i] = -1
		}

		stack = append(stack, arr[i])
		m[arr[i]] = i
	}

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	result := make([]int, n)
	m := make(map[int]int)
	nextLargerElement(nums2, n, m, result)

	output := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		output[i] = result[m[nums1[i]]]
	}

	return output
}
