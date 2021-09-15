package _2_ReverseLinkedListII

import . "learn-leetcode-with-golang/LinkedList"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第一步：从虚拟节点走left-1步，来到left前一个节点
	for i:=0;i<left-1;i++{
		pre = pre.Next
	}
	// 第二步：从pre接着走right-left +1 步来到 right 节点
	rightNode := pre
	for i:=0;i<right-left+1;i++{
		rightNode = rightNode.Next
	}
	// 第三步:切断出子链表，切断连接
	leftNode := pre.Next
	curr := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	// 第四步:反转链表子区间
	reverse(leftNode)
	// 第五步：拼接回来
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func reverse(head *ListNode)  {
	var pre *ListNode
	curr := head
	for curr != nil{
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
}

func reverseBetween2(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	// 需要反转链表前一个节点
	pre := dummyNode
	for i := 0;i < left-1;i++{
		pre = pre.Next
	}
	cur := pre.Next
	// 循环进行插入，将要反转的节点依次插入到cur节点前面，更新pre.Next指向新插入节点
	// 举例： 1 -> 2 -> 3 - > 4
	// 1 为 pre,2 为 cur,next 为 3
	for i := 0;i < right-left;i++{
		// next = 3
		next := cur.Next
		// 2 -> 4
		cur.Next = next.Next
		// 3 -> 2
		next.Next = pre.Next
		// 1 -> 3
		pre.Next = next
		// 插入后结果：1 -> 3 -> 2 -> 4
	}
	return dummyNode.Next
}