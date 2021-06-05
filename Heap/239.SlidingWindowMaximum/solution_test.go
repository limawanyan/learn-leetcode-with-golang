package _39_SlidingWindowMaximum

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(maxSlidingWindow3([]int{1,3,-1,-3,5,3,6},3))
	fmt.Println(maxSlidingWindow4([]int{1,2,3,9,5,6,4},3))
}

func TestList(t *testing.T)  {
	window := []int{1,2,3,4}
	window = window[0:len(window) - 1]
	fmt.Print(window)
}