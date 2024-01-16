/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
    l := 0 // left pointer at the start of the array
    r := len(height)-1 // right pointer at the end of the array
    // store the maxLeft
    maxLeft := height[l]
    // store the maxRight
    maxRight := height[r]
    sum := 0

    for l < r {
        if maxLeft < maxRight {
            l++
            maxLeft = max(height[l], maxLeft)
            sum += maxLeft - height[l]
        } else {
            r--
            maxRight = max(height[r], maxRight)
            sum += maxRight - height[r]
        }
 
    }

    return sum
}
// @lc code=end

