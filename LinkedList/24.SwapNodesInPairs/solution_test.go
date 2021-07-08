package _4_Swap_Nodes_in_pairs

import "testing"

func TestSwapNodesInPairs(t *testing.T) {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	result := swapPairs2(list)
	for result != nil{
		println(result.Val)
		result = result.Next
	}
}