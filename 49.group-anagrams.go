/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, val := range strs {
		var ocrs [26]int
		for _, s := range val {
			ocrs[int(s) - 97]++
		}

		// add to map
		m[ocrs] = append(m[ocrs], val)
	}

	var out [][]string
	for _, value := range m {
		out = append(out, value)
	}

	return out
}
// @lc code=end

