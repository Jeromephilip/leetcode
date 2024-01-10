/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1

	for left < right {

		if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		} else if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		} else {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}

	return true
}

// @lc code=end

