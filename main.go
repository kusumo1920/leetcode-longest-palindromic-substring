package main

import "fmt"

func main() {
	s := "babad"
	output := longestPalindromeSolution4(s)
	fmt.Println(output)
}

func longestPalindromeSolution1(s string) string {
	var resultSubstring string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			currentSubstring := s[i : j+1]
			if isPalindrome(currentSubstring) && len(currentSubstring) > len(resultSubstring) {
				resultSubstring = currentSubstring
			}
		}
	}

	return resultSubstring
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func longestPalindromeSolution2(s string) string {
	check := func(i, j int) bool {
		left := i
		right := j - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true
	}

	for length := len(s); length > 0; length-- {
		for start := 0; start < len(s)-length; start++ {
			if check(start, start+length) {
				return s[start : start+length]
			}
		}
	}

	return ""
}

func longestPalindromeSolution3(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := []int{0, 0}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = []int{i, i + 1}
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff

			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				ans = []int{i, j}
			}
		}
	}

	return s[ans[0] : ans[1]+1]
}

func longestPalindromeSolution4(s string) string {
	expand := func(i, j int) string {
		left, right := i, j

		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		return s[left+1 : right]
	}

	ans := ""

	for i := 0; i < len(s); i++ {
		odd := expand(i, i)
		if len(odd) > len(ans) {
			ans = odd
		}

		even := expand(i, i+1)
		if len(even) > len(ans) {
			ans = even
		}
	}

	return ans
}
