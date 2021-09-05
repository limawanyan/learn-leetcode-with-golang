package _55_MinStack

import "math"

type MinStack2 struct {
	stk  []int
	min  int
	size int
}

/** initialize your data structure here. */
func Constructor2() MinStack2 {
	return MinStack2{stk: make([]int, 0), min: math.MaxInt64}
}

func (this *MinStack2) Push(val int) {
	if this.size == 0 {
		this.min = val
	}
	this.stk = append(this.stk, val-this.min)
	if this.min > val {
		this.min = val
	}
	this.size++
}

func (this *MinStack2) Pop()  {
	if this.stk[this.size-1] < 0 {
		this.min -= this.stk[this.size-1]
	}
	this.stk = this.stk[:this.size-1]
	this.size--
}

func (this *MinStack2) Top() int {
	if this.stk[this.size-1] >= 0 {
		return this.min + this.stk[this.size-1]
	}
	return this.min
}

func (this *MinStack2) GetMin() int {
	return this.min
}