/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
    var m = make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		}
		m[val]++
	}
	return false
}
// @lc code=end

