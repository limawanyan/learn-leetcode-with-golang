package _15_KthLargestElementInAnArray

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(findKthLargest1([]int{3,2,1,5,6,4},2))
	fmt.Println(findKthLargest2([]int{3,2,1,5,6,4},2))
	fmt.Println(findKthLargest3([]int{3,2,1,5,6,4},3))
}