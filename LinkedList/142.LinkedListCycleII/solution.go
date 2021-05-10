package _42_LinkedListCycleII

type ListNode struct {
	Val int
	Next *ListNode
}

// 哈希表法
func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}
	for head != nil{
		if _,ok := m[head];ok{
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 快慢指针
func detectCycle2(head *ListNode) *ListNode {
	// 定义快慢指针
	slow,fast := head,head
	for fast != nil{
		// 慢指针每次走一步
		slow = slow.Next
		// 如果快指针没有下一个节点了，说明没有环
		if fast.Next == nil{
			return nil
		}
		// 快指针每次走两步
		fast = fast.Next.Next
		// 如果快慢指针相遇则代表有环
		if fast == slow {
			// 从头开始遍历，在环入口处与慢指针相遇
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}