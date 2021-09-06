package _34_PalindromeLinkedList

type ListNode struct {
	Val int
	Next *ListNode
}

// 数组+双指针
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	// 遍历一遍节点值添加到数组
	for ;head != nil;head  = head.Next{
		vals = append(vals,head.Val)
	}
	n := len(vals)
	// 双指针从头和尾比较
	for i,v := range vals[:n/2]{
		if v != vals[n-1-i]{
			return false
		}
	}
	return true
}

// 递归
func isPalindrome2(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(node *ListNode) bool
	recursivelyCheck = func(node *ListNode) bool {
		if node != nil{
			if !recursivelyCheck(node.Next){
				return false
			}
			if node.Val != frontPointer.Val{
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}


// 反转链表
func reversList(head *ListNode) *ListNode {
	var prev,cur *ListNode = nil,head
	for cur != nil{
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

// 快慢指针找到前半部分链表尾节点
func endOfFirstHalf(head *ListNode) *ListNode {
	fast,slow := head,head
	// 快指针速度为慢指针两倍，快指针走到尽头，慢指针走一半
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 快慢指针
func isPalindrome3(head *ListNode) bool {
	if head == nil{
		return true
	}
	// 获取前半部分链表尾节点
	firstHalfEnd := endOfFirstHalf(head)
	// 反转后半部分节点
	secondHalfStart := reversList(firstHalfEnd.Next)

	// 与反转后的后半部分链表比较判断是否是回文
	p1,p2 := head,secondHalfStart
	result := true
	for result && p2 != nil{
		if p1.Val != p2.Val{
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 还原链表反转
	firstHalfEnd.Next = reversList(secondHalfStart)
	return result
}


