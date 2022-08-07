package _6_Submissions

type ListNode struct {
	Val  int
	Next *ListNode
}

// Partition 模拟 时间复杂度O(n) 空间复杂度O(1)
func Partition(head *ListNode, x int) *ListNode {
	// 1.模拟法 和两链表合并的题目类似
	// 2.需要维护两个链表，一个链表保存所有小于x值的链，一个链表保存所有大于x值的链
	// 3.最后将两个链表就行合并即可
	// 4.注意利用虚拟头节点指向第一个节点，便于最后结果的返回
	if head == nil || head.Next == nil {
		return head
	}
	dummyMin, dummyMax := &ListNode{}, &ListNode{}
	min, max := dummyMin, dummyMax
	for head != nil {
		if head.Val >= x {
			max.Next = head
			max = max.Next
		} else {
			min.Next = head
			min = min.Next
		}
		head = head.Next
	}
	// 注意这一步，需要将大于x值的链表尾节点Next值设置为nil
	max.Next = nil
	min.Next = dummyMax.Next
	return dummyMin.Next
}
