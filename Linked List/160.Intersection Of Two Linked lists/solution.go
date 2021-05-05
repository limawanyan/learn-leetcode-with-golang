package _60_Intersection_Of_Two_Linked_lists

type ListNode struct {
	Val int
	Next *ListNode
}

// 计算链表长度
func length(head *ListNode) (i int){
	for head != nil{
		i++
		head = head.Next
	}
	return
}

// 对齐起点遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 获取两链表长度
	la,lb := length(headA),length(headB)
	long,short,diff := headA,headB,la - lb
	if la < lb {
		long,short,diff = headB,headA,lb - la
	}
	// 长链表先走，使其与短链表处于同起点
	for i := 0;i < diff;i++{
		long = long.Next
	}
	// 两链表同时走
	for long != nil{
		if long == short{
			return long
		}
		long = long.Next
		short = short.Next
	}
	return nil
}

// 指针从头遍历
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}
	l1,l2 := headA,headB
	for l1 != l2{
		l1 = l1.Next
		if l1 == nil{
			l1 = headB
		}
		l2 = l2.Next
		if l2 == nil {
			l2 = headA
		}
	}
	return l1
}

// 暴力破解法
func getIntersectionNode3(headA,headB *ListNode) *ListNode{
	l1,l2 := headA,headB
	for l1 != nil{
		for l2 != nil {
			if l1 == l2{
				return l1
			}
			l2 = l2.Next
		}
		l1 = l1.Next
		l2 = headB
	}
	return nil
}

// 哈希表法
func getIntersectionNode4(headA,headB *ListNode) *ListNode{
	m := map[*ListNode]struct{}{}
	for headA != nil{
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil{
		if _,ok := m[headB];ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}