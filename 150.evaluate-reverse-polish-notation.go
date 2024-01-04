/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
    stack := make([]int, 0)

	for _, val := range tokens {
		if num, err := strconv.Atoi(val); err == nil { // checks if the value is an integer
			stack = append(stack, num)
		} else {
			// the value is not an integer, pop the top two elements
			secondOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0
			switch val {
				case "+":
					result = firstOperand + secondOperand
				case "-":
					result = firstOperand - secondOperand
				case "*":
					result = firstOperand * secondOperand
				case "/":
					result = firstOperand / secondOperand
			}

			stack = append(stack, result)
		}
	}

	return stack[0]
}
// @lc code=end

