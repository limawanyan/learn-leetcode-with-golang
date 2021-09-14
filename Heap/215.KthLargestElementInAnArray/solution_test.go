package _15_KthLargestElementInAnArray

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	fmt.Println(findKthLargest1([]int{3,2,1,5,6,4},2))
	fmt.Println(findKthLargest2([]int{3,2,1,5,6,4},2))
	fmt.Println(findKthLargest3([]int{3,2,1,5,6,4},3))
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := 5
	l := 3
	fmt.Println(rand.Int() % (r - l + 1))
	fmt.Println(rand.Int() % (r - l + 1))
	fmt.Println(rand.Int() % (r - l + 1))

}