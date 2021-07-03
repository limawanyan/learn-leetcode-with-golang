package DoubleLinkedList

import (
	"fmt"
	"sync"
)

// Golang实现双向链表
// 节点数据
type DoubleObject interface {

}

// 双链表节点
type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}

// 双链表
type DoubleList struct {
	mutex *sync.RWMutex
	Size uint
	Head *DoubleNode
	Tail *DoubleNode
}

// 初始化链表
func (list *DoubleList)Init()  {
	list.mutex = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

// 查询节点
func (list *DoubleList)Get(index uint) *DoubleNode {
	// 判断获取的索引是否在合法范围
	if list.Size == 0 || index > list.Size - 1{
		return nil
	}
	if index == 0{
		return list.Head
	}
	if index == list.Size-1{
		return list.Tail
	}
	node := list.Head
	var i uint
	for i = 1;i <= index;i++{
		node = node.Next
	}
	return node
}

// 追加节点
func (list *DoubleList) Append(node *DoubleNode) bool {
	if node == nil{
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0{
		list.Head = node
		list.Tail = node
		node.Next = nil
		node.Prev = nil
	}else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}

// 指定位置插入节点
func (list *DoubleList)Insert(index uint,node *DoubleNode) bool {
	// 判断插入位置是否超出现有索引
	if index > list.Size || node == nil{
		return false
	}
	// 最后位置插入
	if index == list.Size{
		return list.Append(node)
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		node.Next = list.Head
		list.Head.Prev = node
		list.Head = node
		list.Head.Prev = nil
	}else{
		nextNode := list.Get(index)
		node.Prev = nextNode.Prev
		node.Next = nextNode
		nextNode.Prev.Next = node
		nextNode.Prev = node
	}
	list.Size++
	return true
}

// 删除指定位置节点
func (list *DoubleList) Delete(index uint) bool {
	if index > list.Size - 1{
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	// 删除第一个元素
	if index == 0{
		if list.Size == 1{
			list.Head = nil
			list.Tail = nil
		}else{
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	// 删除最后一个元素
	if index == list.Size - 1{
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}

	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

// 打印链表信息
func (list *DoubleList)Display()  {
	if list == nil || list.Size == 0{
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n",list.Size)
	ptr := list.Head
	for ptr != nil{
		fmt.Printf("data is %v\n",ptr.Data)
		ptr = ptr.Next
	}
}

// 倒序打印链表信息
func (list *DoubleList)Reverse()  {
	if list == nil || list.Size == 0{
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n",list.Size)
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v \n", ptr.Data)
	}
}