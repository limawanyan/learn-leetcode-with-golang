package _8_MergeSortedArray

import (
	"fmt"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	list1 := []int{1,2,3,0,0,0}
	list2 := []int{2,5,6}
	merge(list1,3,list2,3)
	fmt.Println(list1)
}