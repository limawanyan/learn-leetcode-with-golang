package _1_Merge_Two_Sorted_Lists

type ListNode struct {
	Val int
	Next *ListNode
}

// 迭代
func  mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵节点
	prehead := &ListNode{}
	// 保存链表头指针位置
	result := prehead
	// 当两链表不为空的时候循环进行比较
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		}else{
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	// 循环比较后如果还存在剩余元素，指针直接连接剩余有序链表
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}

// 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val < l2.Val{
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}