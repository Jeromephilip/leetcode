/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	// first we sort nums
	sort.Ints(nums)
    m := map[[3]int]bool{}

    for i := 0; i < len(nums)-1; i++ {
        target := -1 * nums[i]
        j := i + 1
        k := len(nums)-1
        for j < k {
            if nums[j]+nums[k]<target {
                j += 1
            } else if nums[j]+nums[k]>target {
                k -= 1
            } else if nums[j] + nums[k] == target {
                arr := [3]int{nums[j], nums[k], nums[i]}
                if _, exists := m[arr]; !exists {
                    m[arr] = true
                }

                j += 1
                k -= 1
            }
        }
    }

    out := [][]int{}
    for key, _ := range m {
        temp := []int{}
        for _, val := range key {
            temp = append(temp, val)
        }
    
        out = append(out, temp)
    }

    return out
}

// @lc code=end

