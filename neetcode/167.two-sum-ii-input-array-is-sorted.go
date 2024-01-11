/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// since the array is sorted, we can use two pointer approach to solve this problem
	// set left pointer to the first index
	left := 0
	right := len(numbers) - 1
	out := make([]int, 0)
	// set right pointer to the second index

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target { // increase the sum
			left++
		} else if sum > target { // decrease the sum
			right--
		} else {
			out = append(out, left + 1)
			out = append(out, right + 1)
			break
		}
	}
	// wrong solution -> doesn't take into account negative numbers
	// for numbers[left] <= target {
	// 	needed := 0
	// 	if (target >= 0 && numbers[left] < 0) {
	// 		needed = target + numbers[left]
	// 	} else {
	// 		needed = target - numbers[left]
	// 	}
		
	// 	if needed == numbers[right] {
	// 		out = append(out, left+1)
	// 		out = append(out, right+1)
	// 		break
	// 	} else if needed > numbers[right] {
	// 		right++
	// 	} else {
	// 		left = right + 1
	// 		right += 2
	// 	}
	// }
	return out
}

// @lc code=end

