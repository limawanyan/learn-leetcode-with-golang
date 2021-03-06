package _5_ReverseNodesInKGroup

type ListNode struct {
	Val int
	Next *ListNode
}

// 模拟
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 头的上一个节点，保存头节点
	hair := &ListNode{Next:head}
	// 分组开始节点
	pre := hair

	for head != nil{
		// 分组结束节点
		tail := pre
		for i := 0;i < k;i++{
			tail = tail.Next
			if tail == nil{
				return hair.Next
			}
		}
		// 保存组结束节点后一个节点
		nex := tail.Next
		// 对分组链表进行反转
		head,tail = myReverse(head,tail)
		// 将头指针的上一个指针指向 分组反转 后的头
		pre.Next = head
		// 将分组反转后的尾指针指向原链表
		tail.Next = nex
		// 将分组反转后的尾指针变成下一分组头指针的上一指针
		pre = tail
		// 将分组尾指针的后一个指针作为下组头指针
		head = nex
	}
	return hair.Next
}

// 链表反转方法
func myReverse(head,tail *ListNode) (*ListNode,*ListNode){
	// 将尾节点指向的节点作为上一个节点
	prev := tail.Next
	// 当前节点
	p := head
	for prev != tail{
		// 下一个节点
		nex := p.Next
		// 将当前节点指向上一个节点
		p.Next = prev
		// 将当前节点作为上一个节点
		prev = p
		// 将下一个节点作为当前节点
		p = nex
	}
	return tail,head
}

// 递归
func reverseKGroup2(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0;i < k;i ++ {
		if tail == nil{
			return head
		}
		// 将指针移动到下一个k group的头指针
		tail = tail.Next
	}
	// 翻转下一个k group
	// 递归返回结果，作为pre，和当前翻转链表连接
	pre := reverseKGroup(tail,k)
	// 反转当前k group
	for i := 0;i < k;i++{
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

//  模拟
func reverseKGroup3(head *ListNode, k int) *ListNode {
	// 保存头节点
	hair := &ListNode{Next: head}
	// 上一个分组尾节点
	pre := hair

	for head != nil{
		// 当前分组尾节点
		tail := pre
		// 进行K分组
		for i := 0;i < k;i++{
			tail = tail.Next
			// 如果不足K个元素,直接返回
			if tail == nil{
				return hair.Next
			}
		}
		// 保存分组后尾节点的后一个节点
		nex := tail.Next
		// 对分组链表进行反转
		head,tail = myReverse2(head,tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = nex
	}
	return hair.Next
}

func myReverse2(head, tail *ListNode) (*ListNode,*ListNode) {
	// 将尾节点下一个节点（下一组的头节点）作为上一个节点
	prev := tail.Next
	// 当前节点
	curr := head
	// 上一个节点再次等于尾节点结束循环
	for prev != tail{
		// 下一个节点
		next := curr.Next
		// 当前节点指向上一个节点
		curr.Next = prev
		// 当前节点作为上一个节点
		prev = curr
		// 下一个节点作为当前节点
		curr = next
	}
	// 尾节点成为反转后组头节点，头节点变为反转后组尾节点
	return tail,head
}