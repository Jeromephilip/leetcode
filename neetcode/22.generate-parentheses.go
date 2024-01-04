/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func decisionTree(n, open, closed int, stack []string, result *[]string) {
	if n == open && n == closed {
		*result = append(*result, strings.Join(stack, ""))
		return
	}

	if open < n {
		stack = append(stack, "(")
		decisionTree(n, open+1, closed, stack, result)
		stack = stack[:len(stack)-1]
	}

	if closed < open {
		stack = append(stack, ")")
		decisionTree(n, open, closed+1, stack, result)
		stack = stack[:len(stack)-1]
	}
}

func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	result := make([]string, 0)

	decisionTree(n, 0, 0, stack, &result)

	return result
}

// @lc code=end

