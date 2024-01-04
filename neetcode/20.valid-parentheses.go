/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if len(stack) == 0 {
			stack = append(stack, char)
		} else if (char == ')' && stack[len(stack)-1] == '(') || (char == '}' && stack[len(stack)-1] == '{') || (char == ']' && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

// @lc code=end

