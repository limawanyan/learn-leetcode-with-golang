package _48_SortList

import . "learn-leetcode-with-golang/LinkedList"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针
	slow,fast := head,head
	// 保存slow前一个节点
	var preSlow *ListNode
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 将链表断开，分成两链
	preSlow.Next = nil
	// 左链
	l := sortList(head)
	// 由链
	r := sortList(slow)
	// 合并左右链
	return mergeList(l,r)
}

func mergeList(l1,l2 *ListNode) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{Val: 0}
	prev := dummy
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		}else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil{
		prev.Next = l1
	}
	if l2 != nil{
		prev.Next = l2
	}
	return dummy.Next
}

// 快速排序
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	quick_sort(head,nil)
	return head
}

func quick_sort(l,r *ListNode) {
	if l == r || l.Next == r {
		return
	}
	mid := getIndex(l,r)
	quick_sort(l,mid)
	quick_sort(mid.Next,r)
}

func getIndex(l, r *ListNode) *ListNode {
	if l == r || l.Next == r {
		return l
	}
	tmp := l.Val
	p1 := l
	p2 := l.Next
	for p2 != r {
		if p2.Val < tmp {
			p1 = p1.Next
			p1.Val,p2.Val = p2.Val,p1.Val
		}
		p2 = p2.Next
	}
	p1.Val,l.Val = l.Val,p1.Val
	return p1
}