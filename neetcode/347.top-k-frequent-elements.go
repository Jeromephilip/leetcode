/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// observation: we know the max frequency is len(nums)
	// create a hashmap of number and frequency:
	freqMap := make(map[int]int)
	for _, val := range nums {
		freqMap[val]++;
	}

	// create a fixed array of size len(nums)
	bucketList := make([][]int, len(nums))
	for key, freq := range freqMap {
		bucketList[freq-1] = append(bucketList[freq-1], key) // appending to nested list
	}

	out := make([]int, 0)
	for i := len(bucketList) - 1; i >= 0; i-- {
		for j:=0; j<len(bucketList[i]); j++ {
			if (k == 0) {
				break
			}
			out = append(out, bucketList[i][j])
			k--
		}
		if (k == 0) {
			break
		}
	}

	return out
}
// @lc code=end

