package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4, 5}))
}

// No.9 Palindrome Number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// No.1480 Running Sum of 1d Array
func runningSum(nums []int) []int {
	var result []int
	last := 0
	for _, v := range nums {
		result = append(result, last+v)
		last += v
	}
	return result
}
