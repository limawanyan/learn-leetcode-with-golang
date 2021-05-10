package _32_ImplementQueueUsingStacks

type MyQueue struct {
	inStack,outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

// 向输入栈添加元素
func (q *MyQueue) Push(x int)  {
	q.inStack = append(q.inStack,x)
}

// 将输入栈所有元素放入输出栈
func (q *MyQueue) in2out()  {
	for len(q.inStack) > 0{
		q.outStack = append(q.outStack,q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

// 取出出队列第一个元素
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0{
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

// 查看出队列第一个元素
func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0{
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

// 判断队列中是否为空
func (q *MyQueue)Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}