/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {

    low := 0
    high := len(nums)-1

    for low <= high {
        mid := low + ((high-low)/2)
        if nums[mid] == target {
            return mid
        }

        if target < nums[mid] {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}
// @lc code=end

