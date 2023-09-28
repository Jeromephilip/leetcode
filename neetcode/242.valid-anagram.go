/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
    var sArray [26]int
	var tArray [26]int
	for i := 0; i < len(s); i++ {
		sArray[int(s[i]) - 97]++;
		tArray[int(t[i]) - 97]++;
	}

	for i:= 0; i < 26; i++ {
		if sArray[i] != tArray[i] {
			return false
		}
	}

	return true
	// PASSED!
}
// @lc code=end

