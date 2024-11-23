/*
	Leetcode Problem: https://leetcode.com/problems/next-greater-element-ii
*/

package main

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	stack := []int{}
	n := len(nums)

	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i%n] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = nums[stack[len(stack)-1]]
		}
		stack = append(stack, i%n)
	}
	return res
}
