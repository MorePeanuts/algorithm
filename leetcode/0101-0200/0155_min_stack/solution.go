// Package leetcode0155 solves LeetCode 155. Min Stack
package leetcode0155

type MinStack struct {
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	top := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	if this.mins[len(this.mins)-1] == top {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
