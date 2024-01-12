/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}


func groupAnagrams(strs []string) [][]string {
    // create a map where sorted_word -> [...words]
	strMap := map[string][]string{}

	for _, val := range strs {
		sorted := sortString(val)
		if _, exists := strMap[sorted]; exists {
			strMap[sorted] = append(strMap[sorted], val)
		} else {
			anagrams := []string{val}
			strMap[sorted] = anagrams
		}
	}

	out := [][]string{}

	// loop again and return the arrays

	for _, val := range strMap {
		out = append(out, val)
	}
	
	return out
}
// @lc code=end

