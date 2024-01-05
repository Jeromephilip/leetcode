/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
type Pair struct {
	index int
	val int
}

func dailyTemperatures(temperatures []int) []int {
    stack := make([]Pair, 0)
	out := make([]int, len(temperatures))

	for index, val := range temperatures {
		if (len(stack) == 0) || (stack[len(stack)-1].val >= val) {
			stack = append(stack, Pair{index, val})
		} else {
			for len(stack) > 0 && stack[len(stack)-1].val < val {
				// pop from stack -> add to count -> append count to out 
				indexOfPoppedElement := stack[len(stack)-1].index
				stack = stack[:len(stack)-1]
				out[indexOfPoppedElement] = index - indexOfPoppedElement
			}
			stack = append(stack, Pair{index, val})
		}
	}


	for len(stack) > 0 {
		index := stack[len(stack)-1].index
		stack = stack[:len(stack)-1]
		out[index] = 0
	}

	return out
}
// @lc code=end

