package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aab", "baa"))
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

// No.1672 Richest Customer Wealth
func maximumWealth(accounts [][]int) int {
	maximum := 0
	for _, account := range accounts {
		sum := 0
		for _, v := range account {
			sum += v
			maximum = max(maximum, sum)
		}
	}
	return maximum
}

// No.412 Fizz Buzz
func fizzBuzz(n int) []string {
	ret := []string{}
	for i := 1; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			ret = append(ret, "FizzBuzz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}

// No.1342 Number of Steps to Reduce a Number to Zero
func numberOfSteps(num int) int {
	step := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		step++
	}
	return step
}

// No.876 Middle of the Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// No.383  Ransom Note
func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, c := range magazine {
		magazineMap[c]++
	}

	for _, c := range ransomNote {
		_, ok := magazineMap[c]
		if !ok {
			return false
		}
		if magazineMap[c] == 0 {
			return false
		}
		magazineMap[c]--
	}
	return true
}
