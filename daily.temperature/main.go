/*
	Leetcode Problem: https://leetcode.com/problems/daily-temperatures/
*/

package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	var s []int

	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[i] >= temperatures[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			result[i] = s[len(s)-1] - i
		} else {
			result[i] = 0
		}

		s = append(s, i)
	}

	return result
}
