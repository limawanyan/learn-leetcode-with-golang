package DoubleLinkedList

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	list := &DoubleList{}
	list.Init()
	list.Append(&DoubleNode{Data: "1"})
	list.Append(&DoubleNode{Data: "2"})
	list.Append(&DoubleNode{Data: "3"})


	list.Insert(0,&DoubleNode{Data: "4"})

	list.Delete(3)

	list.Display()



}
