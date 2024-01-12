/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

type Pair struct {
	firstVal    int
	secondVal   int
}

func threeSum(nums []int) [][]int {
	// calculate the combinations of all pairs possible with indices
	pairMap := map[Pair]int{}

	out := [][]int{}

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {

			p := Pair{0, 0}

			if (nums[i] < nums[j]) {
				p = Pair{nums[i], nums[j]}
			} else {
				p = Pair{nums[j], nums[i]}
			}

			if _, exists := pairMap[p]; exists != true {
				diff := 0 - (nums[i] + nums[j])
				pairMap[p] = diff
			}
		}
	}

	fmt.Println(pairMap)

	for pair, diff := range pairMap {
		for _, val := range nums {
			if val == diff {
				combination := []int{val, pair.firstVal, pair.secondVal}
				out = append(out, combination)
			}
		}
	}

	return out
}

// @lc code=end

