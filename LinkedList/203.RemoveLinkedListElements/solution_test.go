package _03_RemoveLinkedListElements

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	head := &ListNode{Val: 6}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 6}
	list := removeElements2(head,6)
	list.Display()
}

// 输出链表
func (list *ListNode)Display(){
	var res []int
	for list != nil{
		res = append(res,list.Val)
		list = list.Next
	}
	fmt.Println(res)
}