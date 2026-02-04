package leetcode0155

type MinStack2 struct {
	diff []int
	min  int
}

func Constructor2() MinStack2 {
	return MinStack2{
		make([]int, 0),
		0,
	}
}

func (this *MinStack2) Push(val int) {
	if len(this.diff) == 0 {
		this.min = val
	}
	diff := val - this.min
	this.diff = append(this.diff, diff)
	if diff < 0 {
		this.min += diff
	}
}

func (this *MinStack2) Pop() {
	top := this.diff[len(this.diff)-1]
	if top < 0 {
		this.min -= top
	}
	this.diff = this.diff[:len(this.diff)-1]
}

func (this *MinStack2) Top() int {
	top := this.diff[len(this.diff)-1]
	if top < 0 {
		return this.min
	} else {
		return this.min + top
	}
}

func (this *MinStack2) GetMin() int {
	return this.min
}
