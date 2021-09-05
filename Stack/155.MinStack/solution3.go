package _55_MinStack

type MinStack3 struct {
	s []Node
}

type Node struct {
	d int //data
	m int //current min
}

/** initialize your data structure here. */
func Constructor3() MinStack3 {
	return MinStack3{}
}

func (this *MinStack3) Push(x int)  {
	d := Node{d:x,m:x}
	if len(this.s)>0 &&this.s[len(this.s)-1].m< x{
		d.m=this.s[len(this.s)-1].m
	}
	this.s = append(this.s,d)
}

func (this *MinStack3) Pop()  {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack3) Top() int {
	return this.s[len(this.s)-1].d
}

func (this *MinStack3) GetMin() int {
	return this.s[len(this.s)-1].m
}