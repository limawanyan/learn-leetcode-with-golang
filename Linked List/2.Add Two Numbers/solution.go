package __Add_Two_Numbers

type  ListNode struct {
	Val int
	Next *ListNode
}

// 递归法
func  addTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode  {
	// 当两个链表都为空时结束
	if l1 == nil && l2 == nil{
		return nil
	}
	// 如果l1链表为空，直接返回链表2节点
	if l1 == nil {
		return l2
	}
	// 如果l2链表为空，直接返回链表1节点
	if l2 == nil{
		return l1
	}
	// 计算两链表节点和值
	sum := l1.Val + l2.Val
	// 递归计算节点和值
	nextNode := addTwoNumbers(l1.Next,l2.Next)
	// 如果值小于10代表没有进位，返回一个节点，值为sum，next指针指向递归处理好的节点
	if sum < 10 {
		return &ListNode{sum,nextNode}
	}else{
		// 存在进位，生成一个进位节点，值为1。因为最大两数相加9 + 9 = 18 进位为1
		tempNode := &ListNode{
			1,
			nil,
		}
		// 返回节点，值为 sum - 10，Next为 上一步递归计算好的值 + 进位节点
		return &ListNode{
			sum - 10,
			addTwoNumbers(nextNode,tempNode),
		}
	}
}

// 模拟
func addTwoNumbers2(l1,l2 *ListNode) (head *ListNode){
	// 定义一个链表
	var tail *ListNode
	// 保存进位
	carry := 0
	// 当两个链表都为空的时候结束循环
	for l1 != nil || l2 != nil{
		// 初始化两链表节点值为0
		n1,n2 := 0,0
		// 如果节点值不为空，给上面变量赋值，然后将链表指向下一个节点
		if l1 != nil{
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			n2 = l2.Val
			l2 = l2.Next
		}
		// 两节点和值 + 进位
		sum := n1 + n2 + carry
		// 计算和值和进位
		sum,carry = sum %10,sum /10
		// 初始化处理，顺便赋返回值
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		}else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 如果存在进位再附加一个节点
	if carry > 0{
		tail.Next = &ListNode{Val: carry}
	}
	return
}