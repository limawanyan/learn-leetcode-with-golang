package _32_ImplementQueueUsingStacks

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}