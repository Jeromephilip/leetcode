/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
    maxArea := 0

    l := 0
    r := len(height)-1

    for l < r {
        length := min(height[l], height[r])
        width := r - l
        area := length * width

        if (maxArea < area) {
            maxArea = area
        }

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return maxArea
}
// @lc code=end

