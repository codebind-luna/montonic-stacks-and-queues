/*
Leetcode Problem: https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop
*/
package main

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	s := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && prices[i] < prices[s[len(s)-1]] {
			s = s[:len(s)-1]
		}

		if len(s) > 0 {
			result[i] = prices[i] - prices[s[len(s)-1]]
		} else {
			result[i] = prices[i]
		}

		s = append(s, i)

	}
	return result
}
