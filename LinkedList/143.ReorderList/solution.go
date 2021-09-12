package _43_ReorderList

type ListNode struct {
	Val int
	Next *ListNode
}

// 线性表 时间空间复杂度O（n）
func reorderList(head *ListNode)  {
	if head == nil{
		return
	}
	// 将所有节点存储到数组中
	nodes := []*ListNode{}
	for node := head;node != nil;node = node.Next{
		nodes = append(nodes,node)
	}
	// 定义两指针
	i,j := 0,len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next= nil
}

// 寻找链表中点 + 后半部分链表反转 + 间隔合并两链表
// 时间复杂度O（n）空间复杂度O（1）
func reorderList2(head *ListNode) {
	if head == nil{
		return
	}
	// 找到中间节点
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	// 中间节点指向nil分割两链表
	mid.Next = nil
	// 反转后半部分链表
	l2 = reverseList(l2)
	// 合并两链表
	mergeList(l1,l2)
}

// 找链表中点
func middleNode(head *ListNode) *ListNode {
	slow,fast := head,head
	// 利用快慢指针找到上半部分链表尾节点
	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev,cur *ListNode = nil,head
	for cur !=  nil{
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

// 合并链表
func mergeList(l1, l2 *ListNode) {
	var l1Tmp,l2Tmp *ListNode
	for l1 != nil && l2 != nil{
		l1Tmp = l1.Next
		l2Tmp = l2.Next
		l1.Next = l2
		l1 = l1Tmp
		l2.Next = l1
		l2 = l2Tmp
	}
}