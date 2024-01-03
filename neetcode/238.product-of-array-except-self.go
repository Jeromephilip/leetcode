/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	prefixMultiplier := 1
	postfixMultiplier := 1
	// populating prefix values
	for key, _ := range nums {
		prefixMultiplier *= nums[key]
		postfixMultiplier *= nums[len(nums)-1-key]
		prefix[key] = prefixMultiplier
		postfix[len(nums)-1-key] = postfixMultiplier
	}

	out := make([]int, 0)

	for key, _ := range nums {
		if key == 0 {
			out = append(out, postfix[key+1])
		} else if key == len(nums)-1 {
			out = append(out, prefix[key-1])
		} else {
			out = append(out, prefix[key-1]*postfix[key+1])
		}
	}

	return out
}

// @lc code=end

