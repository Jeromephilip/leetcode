/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	longest := 0

    for _, val := range nums {
		set[val] = true
	}


	for _, val := range nums {
		// check if starting number in a sequence
		if _, exist := set[val-1]; !exist {
			// keep incrementing by one, starting with the leftmost in a sequence 
			// to the rightmost in a sequence in the set
			curVal := val
			curLen := 1
			for true {
				if _, exists := set[curVal + 1]; !exists {
					break
				}
				curVal++
				curLen++
			}
			longest = int(math.Max(float64(curLen), float64(longest)))
		}
	}

	return longest
}
// @lc code=end

