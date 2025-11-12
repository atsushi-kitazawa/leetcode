package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	validPhoneNumber()
}

// No.193 Valid Phone number
func validPhoneNumber() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("failed to open file.txt")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	re1 := regexp.MustCompile(`^\(\d{3}\) \d{3}-\d{4}$`)
	re2 := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
	for scanner.Scan() {
		line := scanner.Text()
		if re1.MatchString(line) || re2.MatchString(line) {
			fmt.Println(line)
		}
	}
}

// No.112 Path Sum
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return (targetSum - root.Val) == 0
	}

	targetSum -= root.Val

	if hasPathSum(root.Left, targetSum) {
		return true
	}

	if hasPathSum(root.Right, targetSum) {
		return true
	}

	return false
}

// No.14 Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	rPrefix := []rune(prefix)
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		rStr := []rune(strs[i])
		for j, c := range rStr {
			if j >= len(rPrefix) {
				break
			}
			if rPrefix[j] != c {
				rPrefix = rPrefix[:j]
				break
			}
			if j == len(rStr)-1 {
				rPrefix = rPrefix[:j+1]
				break
			}
		}
	}
	return string(rPrefix)
}

// No.169 Majority Element
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	for _, v := range nums {
		cnt, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = cnt + 1
			if m[v] > len(nums)/2 {
				return v
			}
		}
	}
	return 0
}

// No.415 Add Strings
func addStrings(num1 string, num2 string) string {
	// fmt.Printf("start %s %s\n", num1, num2)
	// sum := make([]string, 0)
	// rnum1 := []rune(num1)
	// rnum2 := []rune(num2)
	// carryover := false
	// for i, j := len(rnum1)-1, len(rnum2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
	// 	var n1, n2 int
	// 	if i >= 0 {
	// 		// fmt.Printf("num1:%s\n", string(rnum1[i]))
	// 		n1, _ = strconv.Atoi(string(rnum1[i]))
	// 	}
	// 	if j >= 0 {
	// 		// fmt.Printf("num2:%s\n", string(rnum2[j]))
	// 		n2, _ = strconv.Atoi(string(rnum2[j]))
	// 	}
	// 	n := n1 + n2
	// 	if carryover {
	// 		n++
	// 	}
	// 	if len(strconv.Itoa(n)) == 2 {
	// 		carryover = true
	// 	} else {
	// 		carryover = false
	// 	}

	// 	fmt.Println(n)
	// }
	return ""
}

// No.58 Length of Last Word
func lengthOfLastWord(s string) int {
	slice := strings.Fields(s)
	return len([]rune(slice[len(slice)-1]))
}

// No.35 Search Insert Position
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			return i - 0
		}
	}
	return len(nums)
}

// No.268 Missing Number
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

// No.26 Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		if previous == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		previous = nums[i]
	}
	return len(nums)
}

// No.27 Remove Element
func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			count++
			i--
			continue
		}
	}
	return len(nums)
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

// No.283 Move Zeros
// you must do this in-place without making a copy of the array
func moveZeros(nums []int) {
	loopCount := 0
	for i := 0; i < len(nums); i++ {
		if loopCount == len(nums) {
			break
		}

		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i--
		}
		loopCount++
	}

	fmt.Println(nums)
}
