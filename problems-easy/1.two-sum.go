/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
        if pos, ok := m[target-nums[i]]; ok {
            return []int{pos, i}
        } else {
			m[nums[i]] = i
		}
	}
	return []int{}
}
// @lc code=end

