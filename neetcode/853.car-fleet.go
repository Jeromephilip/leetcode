/*
 * @lc app=leetcode id=853 lang=golang
 *
 * [853] Car Fleet
 */

// @lc code=start

type Pair struct {
	position int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
    // sort the array such position wise (nlogn)
	arr := make([]Pair, 0)
	stack := make([]float64, 0)

	for i := 0; i < len(position); i++ {
		pair := Pair{position[i], speed[i]}
		arr = append(arr, pair)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].position < arr[j].position
	})
	
	for i := len(arr)-1; i >= 0; i-- {
		time := float64(target - arr[i].position) / float64(arr[i].speed)
		stack = append(stack, time)

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
// @lc code=end

