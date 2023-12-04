/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
    // prefix array
	var prefix []int
	// postfix array
	var suffix []int

	// calculate all the products previous to nums[i]
	prefixProduct := 1
	for _, val := range nums {
		prefixProduct = prefixProduct * val
		prefix = append(prefix,	prefixProduct)
	}

	suffixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		suffixProduct = suffixProduct * nums[i]
		suffix = append(suffix, suffixProduct)
	}

	var out []int
	
}
// @lc code=end

