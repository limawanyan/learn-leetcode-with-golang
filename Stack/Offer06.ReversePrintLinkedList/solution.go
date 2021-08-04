package Offer06_ReversePrintLinkedList

type ListNode struct {
	     Val int
	     Next *ListNode
}
// 栈  时间复杂度O（n） 空间复杂度O（n）
func reversePrint(head *ListNode) (res []int) {
	var stack []*ListNode
	// 压入栈
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	// 弹出栈
	for len(stack) != 0 {
		res = append(res, stack[len(stack)-1].Val)
		stack = stack[:len(stack)-1]
	}
	return res
}

// 遍历 时间复杂度O(N) 空间复杂度O(1)
func reversePrint2(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append([]int{head.Val}, ans...)
		head = head.Next
	}
	return ans
}

// 递归
func reversePrint3(head *ListNode) []int {
	if head == nil{
		return nil
	}
	return append(reversePrint3(head.Next),head.Val)
}

// 反转链表
func reversePrint4(head *ListNode) (res []int) {
	if head == nil{
		return nil
	}
	// 反转链表
	var pre *ListNode
	for head != nil{
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	for pre != nil{
		res = append(res,pre.Val)
		pre = pre.Next
	}
	return res
}

// 反转数组
func reversePrint5(head *ListNode) (res []int) {
	if head == nil{
		return nil
	}
	// 顺序存储结果
	for head != nil{
		res = append(res,head.Val)
		head = head.Next
	}
	// 将结果倒序
	for i,j := 0,len(res)-1;i < j;{
		res[i],res[j] = res[j],res[i]
		i++
		j--
	}
	return res
}