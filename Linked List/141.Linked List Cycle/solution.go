package _41_Linked_List_Cycle

type ListNode struct{
	Val int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil{
		slow,fast = slow.Next,fast.Next.Next
		if slow == fast{
			return true
		}
	}
	return false
}

// 哈希表
func hasCycle2(head *ListNode) bool {
	// 通过map实现set
	m := map[*ListNode]struct{}{}
	for head != nil{
		if _,ok := m[head];ok{
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}