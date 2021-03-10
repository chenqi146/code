package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 优化版
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	source := x
	r := 0
	for x != 0 {
		r = r*10 + x%10
		x = x / 10
	}

	return source == r
}

// 最终优化版
func isPalindrome3(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	r := 0
	// 当X小于转换的数字时跳出
	for x > r {
		r = r*10 + x%10
		x = x / 10
	}

	// 如果是奇数   此时   r / 10 == x
	// 如果是偶数  r == x
	return r == x || r/10 == x
}

func main() {
	println(isPalindrome2(2221222))
}
