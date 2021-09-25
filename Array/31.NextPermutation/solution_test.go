package _1_NextPermutation

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	nums := []int{3,2,1}
	nextPermutation(nums)
	fmt.Println(nums)
}