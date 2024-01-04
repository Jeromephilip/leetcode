/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
    stack []int
	minStack []int
}


func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}


func (this *MinStack) Pop()  {
    val := this.stack[len(this.stack)-1]
    if len(this.minStack) != 0 && val == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

