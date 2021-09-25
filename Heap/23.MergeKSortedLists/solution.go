package _3_MergeKSortedLists

import "container/heap"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	result := lists[0]
	for i:=1;i < len(lists);i++{
		result=Merge(result,lists[i])
	}
	return result
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val{
		res = l2
		res.Next = Merge(l1,l2.Next)
	}else {
		res = l1
		res.Next = Merge(l1.Next,l2)
	}
	return res
}

type node_slice []*ListNode

func (n node_slice) Len() int {
	return len(n)
}

func (n node_slice) Less(i,j int) bool {
	return n[i].Val < n[j].Val
}

func (n *node_slice) Swap(i, j int) {
	(*n)[i],(*n)[j] = (*n)[j],(*n)[i]
}

func (n *node_slice) Push(h interface{})  {
	node := h.(*ListNode)
	*n = append(*n,node)
}

func (n *node_slice) Pop() interface{} {
	tmp := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return tmp
}

func mergeKLists2(lists []*ListNode) *ListNode {
	h := &node_slice{}
	heap.Init(h)
	for _,v := range lists{
		if v == nil{
			continue
		}
		heap.Push(h,v)
	}
	ret := &ListNode{}
	node := ret
	for h.Len() > 0{
		ln := heap.Pop(h).(*ListNode)
		node.Next = ln
		if ln.Next != nil{
			heap.Push(h,ln.Next)
		}
		node = node.Next
	}
	return ret.Next
}