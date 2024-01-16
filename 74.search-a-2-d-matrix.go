/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    rl := 0
    rh := len(matrix)-1

    ri := -1 // row index to iterate over

    for rl <= rh {
        rm := rl + ((rh-rl)/2)
        if matrix[rm][0] <= target && matrix[rm][len(matrix[rm])-1] >= target {
            ri = rm
            break
        }

        if matrix[rm][0] > target {
            rh = rm - 1
        } else {
            rl = rm + 1
        }
    }

    if ri < 0 {
        return false
    }

    cl := 0
    ch := len(matrix[ri])-1

    fmt.Println(ri)

    for cl <= ch {
        cm := cl + ((ch - cl)/2)
        
        if matrix[ri][cm] == target {
            return true
        }

        if matrix[ri][cm] > target {
            ch = cm - 1
        } else {
            cl = cm + 1
        }
    }

    return false
}
// @lc code=end

