/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

/* Initial idea is to create a map where the key is an array
containing the occurences of a given character in a string and 
the value is all the strings that match the occurence key. Another
way to solve this is having the key as a sorted string
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

